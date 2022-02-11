package models

import (
	"app/database"
	"sort"
	"time"
)

type Tag struct {
	ID        uint64    `json:"id"`
	Slug      string    `json:"slug" gorm:"index:idx_tag_slug,unique"`
	Title     string    `json:"title" gorm:"index:idx_tag_title,unique"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TagInterface struct {
	ID        uint64    `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Count     int       `json:"count"`
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

func TagOccurrencesInProjects(projects []*Project, tag *Tag) int {
	count := 0
	for _, project := range projects {
		for _, projectTag := range project.Tags {
			if projectTag.ID == tag.ID {
				count = count + 1
			}
		}
	}

	return count
}

func TagsOccurrencesInProjects(projects []*Project) []*TagInterface {
	var tags []*TagInterface
	tagsMap := make(map[string]*TagInterface)
	for _, project := range projects {
		for _, tag := range project.Tags {
			if tagInterface, ok := tagsMap[tag.Slug]; ok {
				tagInterface.Count = tagInterface.Count + 1
			} else {
				tagsMap[tag.Slug] = &TagInterface{
					ID:        tag.ID,
					Slug:      tag.Slug,
					Title:     tag.Title,
					CreatedAt: tag.CreatedAt,
					UpdatedAt: tag.UpdatedAt,
					Count:     1,
				}
			}
		}
	}

	keys := make([]string, 0, len(tagsMap))
	for key := range tagsMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		tags = append(tags, tagsMap[key])
	}

	return tags
}
