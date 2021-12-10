package models

import (
	"app/database"
	"time"
)

type Project struct {
	ID          uint64    `json:"id"`
	Slug        string    `json:"slug" gorm:"index:idx_project_slug,unique"`
	Title       string    `json:"title" gorm:"index:idx_project_title,unique"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	URL         *string   `json:"url"`
	Git         *string   `json:"git"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Tags        []*Tag    `json:"tags" gorm:"many2many:project_tags"`
}

func GetProjects() ([]*Project, error) {
	var projects []*Project
	if db, err := database.Open(); err == nil {
		results := db.Preload("Tags").Find(&projects)
		return projects, results.Error
	} else {
		return nil, err
	}
}

func FindProject(project *Project) (bool, error) {
	return database.RecordExist(&Project{}, "id = ?", project.ID)
}

func UpdateProject(project *Project, id string) error {
	projectRecord, err := GetProject(id)
	if err != nil {
		return err
	}

	projectRecord.Title = project.Title
	projectRecord.Slug = project.Slug
	projectRecord.Description = project.Description
	projectRecord.Image = project.Image
	projectRecord.URL = project.URL
	projectRecord.Git = project.Git
	projectRecord.Tags = project.Tags

	return database.UpdateRecord(&projectRecord)
}

func CreateProject(project *Project) error {
	return database.CreateRecord(project)
}

func GetProject(id string) (*Project, error) {
	var project *Project
	if db, err := database.Open(); err == nil {
		results := db.Preload("Tags").Where("id = ?", id).Find(&project)
		return project, results.Error
	} else {
		return nil, err
	}
}

func DeleteProject(id string) error {
	return database.DeleteRecord(&Project{}, id)
}
