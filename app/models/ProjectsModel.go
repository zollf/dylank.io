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
	found, err := FindProject(project)

	if err != nil {
		return nil
	}

	if found {
		return UpdateProject(project)
	} else {
		return CreateProject(project)
	}
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

func FindProject(project *Project) (bool, error) {
	var found bool = false
	err := database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		count, err := client.Database("db").Collection(ProjCol).CountDocuments(ctx, bson.M{"_id": project.ID})

		if err != nil {
			return err
		}

		found = count != 0
		return nil
	})

	return found, err
}

func UpdateProject(project *Project) error {
	return database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		_, err := client.Database("db").Collection(ProjCol).UpdateOne(ctx, bson.M{"_id": project.ID}, bson.M{"$set": project})
		return err
	})
}

func CreateProject(project *Project) error {
	return database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		_, err := client.Database("db").Collection(ProjCol).InsertOne(ctx, project)
		return err
	})
}

func GetProject(id string) (*Project, error) {
	var project *Project
	err := database.GetMongo(func(ctx context.Context, client *mongo.Client) error {
		pid, pid_err := primitive.ObjectIDFromHex(id)

		if pid_err != nil {
			return pid_err
		}

		result := client.Database("db").Collection(ProjCol).FindOne(ctx, bson.M{"_id": pid})

		if result.Err() != nil {
			return result.Err()
		}

		return result.Decode(&project)
	})

	return project, err
}
