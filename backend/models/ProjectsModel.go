package models

import (
	"app/database"
	"fmt"

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

const ProjCol = "project"

func CreateOrEditProject(project *Project) error {
	if found, err := FindProject(project); err != nil {
		return fmt.Errorf("error occurred when trying to find if project exists")
	} else {
		if found {
			return UpdateProject(project)
		} else {
			return CreateProject(project)
		}
	}
}

func GetProjects() ([]*Project, error) {
	var projects []*Project
	err := database.GetDocuments(ProjCol, func(cur *mongo.Cursor) error {
		var project *Project
		if err := cur.Decode(&project); err != nil {
			return err
		} else {
			projects = append(projects, project)
			return nil
		}
	})

	return projects, err
}

func FindProject(project *Project) (bool, error) {
	return database.DocumentExist(ProjCol, bson.M{"_id": project.ID})
}

func UpdateProject(project *Project) error {
	return database.UpdateDocument(ProjCol, bson.M{"_id": project.ID}, bson.M{"$set": project})
}

func CreateProject(project *Project) error {
	return database.CreateDocument(ProjCol, project)
}

func GetProject(id string) (*Project, error) {
	if pid, pid_err := primitive.ObjectIDFromHex(id); pid_err != nil {
		return nil, pid_err
	} else {
		var project *Project
		err := database.GetDocument(ProjCol, bson.M{"_id": pid}, func(res *mongo.SingleResult) error {
			return res.Decode(&project)
		})
		return project, err
	}
}

func DeleteProject(id string) error {
	if pid, pid_err := primitive.ObjectIDFromHex(id); pid_err != nil {
		return pid_err
	} else {
		return database.DeleteDocument(ProjCol, bson.M{"_id": pid})
	}
}
