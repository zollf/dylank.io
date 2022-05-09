package tags

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

type TagCreateRequest struct {
	Title    string `json:"title" validate:"required"`
	Redirect string `json:"redirect"`
}

type TagEditRequest struct {
	ID       string `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Redirect string `json:"redirect"`
}

type TagDeleteRequest struct {
	ID       string `json:"id" validate:"required"`
	Redirect string `json:"redirect"`
}

type TagGetRequest struct {
	ID string `param:"id" validate:"required"`
}

func All() ([]*Tag, error) {
	var tags []*Tag
	err := database.GetRecords(&tags)
	return tags, err
}

func Exists(cond ...interface{}) (bool, error) {
	return database.RecordExist(&Tag{}, cond)
}

func (tag Tag) Update() error {
	tagRecord, not_found := Find(tag.ID)
	if not_found != nil {
		return not_found
	}

	tagRecord.Title = tag.Title
	tagRecord.Slug = tag.Slug

	return database.UpdateRecord(&tagRecord)
}

func (tag Tag) Create() error {
	return database.CreateRecord(&tag)
}

func Find(id interface{}) (*Tag, error) {
	var tag *Tag
	err := database.GetRecord(&tag, "id = ?", id)
	return tag, err
}

func (tag Tag) Delete() error {
	return database.DeleteRecord(&Tag{}, tag.ID)
}
