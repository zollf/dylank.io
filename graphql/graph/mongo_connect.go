package graph

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type query func(context.Context, *mongo.Client) error

const defaultUri = "mongodb://mongo:27017"

func GetMongo(q query) error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = defaultUri
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	err = q(ctx, client)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
