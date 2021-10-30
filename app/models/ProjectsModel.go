package models

import (
	"app/database"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Project struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Slug        string             `json:"slug" bson:"slug"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Image       string             `json:"image" bson:"image"`
	URL         *string            `json:"url" bson:"url"`
	Git         *string            `json:"git" bson:"git"`
	DateCreated string             `json:"dateCreated" bson:"dateCreated"`
	DateUpdated string             `json:"dateUpdated" bson:"dateUpdated"`
	Tags        []*Tag             `json:"tags" bson:"tags"`
}

type Tag struct {
	ID    string `json:"id" bson:"_id"`
	Slug  string `json:"slug" bson:"slug"`
	Title string `json:"title" bson:"title"`
}

const ProjCol = "project"

func CreateOrEditProject(project *Project) error {
	return nil
}

func GetProjects() ([]*Project, error) {
	var projects []*Project
	err := database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		cur, err := client.Database("db").Collection(ProjCol).Find(ctx, bson.D{})

		if err != nil {
			return err
		}

		for cur.Next(ctx) {
			var result *Project
			err := cur.Decode(&result)

			if err != nil {
				return err
			}

			projects = append(projects, result)
		}

		return nil
	})

	return projects, err
}
