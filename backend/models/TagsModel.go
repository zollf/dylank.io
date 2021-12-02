package models

import (
	"app/database"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tag struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Slug        string             `json:"slug" bson:"slug"`
	Title       string             `json:"title" bson:"title"`
	DateCreated string             `json:"dateCreated" bson:"dateCreated"`
	DateUpdated string             `json:"dateUpdated" bson:"dateUpdated"`
}

const TagCol = "tag"

func CreateOrEditTag(tag *Tag) error {
	if found, err := FindTag(tag); err != nil {
		return fmt.Errorf("error occurred when trying to find if tag exists")
	} else {
		if found {
			return UpdateTag(tag)
		} else {
			tag_exist, tag_err := FindTagByTitle(tag)
			if tag_err != nil {
				return tag_err
			}

			if tag_exist {
				return fmt.Errorf("Tag title already exists")
			}

			return CreateTag(tag)
		}
	}
}

func GetTags() ([]*Tag, error) {
	var tags []*Tag
	err := database.GetDocuments(TagCol, func(cur *mongo.Cursor) error {
		var tag *Tag
		if err := cur.Decode(&tag); err != nil {
			return err
		} else {
			tags = append(tags, tag)
			return nil
		}
	})

	return tags, err
}

func FindTag(tag *Tag) (bool, error) {
	return database.DocumentExist(TagCol, bson.M{"_id": tag.ID})
}

func FindTagByTitle(tag *Tag) (bool, error) {
	return database.DocumentExist(TagCol, bson.M{"title": tag.Title})
}

func UpdateTag(tag *Tag) error {
	return database.UpdateDocument(TagCol, bson.M{"_id": tag.ID}, bson.M{"$set": tag})
}

func CreateTag(tag *Tag) error {
	return database.CreateDocument(TagCol, tag)
}

func GetTag(id string) (*Tag, error) {
	if pid, pid_err := primitive.ObjectIDFromHex(id); pid_err != nil {
		return nil, pid_err
	} else {
		var tag *Tag
		err := database.GetDocument(TagCol, bson.M{"_id": pid}, func(res *mongo.SingleResult) error {
			return res.Decode(&tag)
		})
		return tag, err
	}
}

func DeleteTag(id string) error {
	if pid, pid_err := primitive.ObjectIDFromHex(id); pid_err != nil {
		return pid_err
	} else {
		return database.DeleteDocument(TagCol, bson.M{"_id": pid})
	}
}

// Given a tag, check it exists amongst an array of tags
func CheckTagExistsInTags(needle *Tag, haystack []*Tag) bool {
	for _, needleInHay := range haystack {
		if needleInHay.ID == needle.ID {
			return true
		}
	}
	return false
}
