package database

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultUri = "mongodb://mongo:27017"

// Open a mongo db connection then query.
//
// Usage:
//	err := GetMongo(func(ctx context.Context, client *mongo.Client) error {
//		// do query
//		return err
//	})
//
func GetMongo(q func(context.Context, *mongo.Client) error) error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = defaultUri
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	defer client.Disconnect(ctx)

	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return err
	}

	err = q(ctx, client)
	return err
}

// Gets documents from mongodb.
func GetDocuments(col string, callback func(cur *mongo.Cursor) error) error {
	return GetMongo(func(ctx context.Context, client *mongo.Client) error {
		if cur, err := client.Database("db").Collection(col).Find(ctx, bson.D{}); err != nil {
			return err
		} else {
			for cur.Next(ctx) {
				if err := callback(cur); err != nil {
					return err
				}
			}
			return nil
		}
	})
}

// Checks if document exists within mongodb
//
// Usage:
//	DocumentExist('object_column', bson.M{"_id": object.ID})
func DocumentExist(col string, filter bson.M) (bool, error) {
	var found bool = false
	err := GetMongo(func(ctx context.Context, client *mongo.Client) error {
		if count, err := client.Database("db").Collection(col).CountDocuments(ctx, filter); err != nil {
			return err
		} else {
			found = count != 0
			return nil
		}
	})
	return found, err
}

// Updates one document
//
// Usage:
//	UpdateDocument('object_column', bson.M{"_id": object.ID}, bson.M{"$set": object})
func UpdateDocument(col string, filter bson.M, update bson.M) error {
	return GetMongo(func(ctx context.Context, client *mongo.Client) error {
		_, err := client.Database("db").Collection(col).UpdateOne(ctx, filter, update)
		return err
	})
}

func CreateDocument(col string, data interface{}) error {
	return GetMongo(func(ctx context.Context, client *mongo.Client) error {
		_, err := client.Database("db").Collection(col).InsertOne(ctx, data)
		return err
	})
}

func GetDocument(col string, filter bson.M, callback func(res *mongo.SingleResult) error) error {
	err := GetMongo(func(ctx context.Context, client *mongo.Client) error {
		if result := client.Database("db").Collection(col).FindOne(ctx, filter); result.Err() != nil {
			return result.Err()
		} else {
			return callback(result)
		}
	})
	return err
}

func DeleteDocument(col string, filter bson.M) error {
	return GetMongo(func(ctx context.Context, client *mongo.Client) error {
		res, err := client.Database("db").Collection(col).DeleteOne(ctx, filter)
		if res.DeletedCount == 0 {
			return errors.New("Document does not exist")
		}
		return err
	})
}
