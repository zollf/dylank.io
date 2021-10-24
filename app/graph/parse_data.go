package graph

import (
	"app/database"
	"app/graph/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseProjects() []*model.Project {
	var projects []*model.Project
	err := database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		cur, err := client.Database("db").Collection("projects").Find(ctx, bson.D{{}})
		for cur.Next(ctx) {
			var result *model.Project
			err := cur.Decode(&result)

			if err != nil {
				return err
			}

			projects = append(projects, result)
		}
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	return projects
}
