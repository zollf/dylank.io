package models

import (
	"app/database"
	"time"
)

type Tag struct {
	ID        uint64    `json:"id"`
	Slug      string    `json:"slug" gorm:"index:idx_tag_slug,unique"`
	Title     string    `json:"title" gorm:"index:idx_tag_title,unique"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetTags() ([]*Tag, error) {
	var tags []*Tag
	err := database.GetRecords(&tags)
	return tags, err
}

func FindTag(tag *Tag) (bool, error) {
	return database.RecordExist(&Tag{}, "id = ?", tag.ID)
}

func FindTagByTitle(tag *Tag) (bool, error) {
	return database.RecordExist(&Tag{}, "title = ?", tag.Title)
}

func UpdateTag(tag *Tag, id string) error {
	tagRecord, err := GetTag(id)
	if err != nil {
		return err
	}

	tagRecord.Title = tag.Title
	tagRecord.Slug = tag.Slug

	return database.UpdateRecord(&tagRecord)
}

func CreateTag(tag *Tag) error {
	return database.CreateRecord(tag)
}

func GetTag(id string) (*Tag, error) {
	var tag *Tag
	err := database.GetRecord(&tag, "id = ?", id)
	return tag, err
}

func DeleteTag(id string) error {
	return database.DeleteRecord(&Tag{}, id)
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
