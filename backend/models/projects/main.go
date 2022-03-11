package projects

import (
	"app/database"
	"app/models/assets"
	"app/models/tags"
	"time"
)

type Project struct {
	ID          uint64          `json:"id"`
	Slug        string          `json:"slug" gorm:"index:idx_project_slug,unique"`
	Title       string          `json:"title" gorm:"index:idx_project_title,unique"`
	Description string          `json:"description"`
	Assets      []*assets.Asset `json:"assets" gorm:"many2many:project_assets"`
	URL         *string         `json:"url"`
	Git         *string         `json:"git"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Tags        []*tags.Tag     `json:"tags" gorm:"many2many:project_tags"`
}

// List all projects.
func All() ([]*Project, error) {
	var projects []*Project
	if db, err := database.Open(); err == nil {
		results := db.Preload("Tags").Preload("Assets").Find(&projects)
		return projects, results.Error
	} else {
		return nil, err
	}
}

// Check if project exists.
func Exists(id string) (bool, error) {
	return database.RecordExist(&Project{}, "id = ?", id)
}

// Updates the current project.
func (project Project) Update() error {
	projectRecord, err := Find(project.ID)

	if err != nil {
		return err
	}

	// !TODO: fix
	projectRecord.Title = project.Title
	projectRecord.Slug = project.Slug
	projectRecord.Description = project.Description
	projectRecord.URL = project.URL
	projectRecord.Git = project.Git

	if db, err := database.Open(); err == nil {
		db.Save(&projectRecord)
		db.Model(&projectRecord).Association("Tags").Replace(project.Tags)
		db.Model(&projectRecord).Association("Assets").Replace(project.Assets)
		return nil
	} else {
		return err
	}
}

// Creates new project.
func (project Project) Create() error {
	return database.CreateRecord(&project)
}

func Find(id interface{}) (*Project, error) {
	var project *Project
	if db, err := database.Open(); err == nil {
		results := db.
			Preload("Tags").
			Preload("Assets").
			Where("id = ?", id).
			Find(&project)
		return project, results.Error
	} else {
		return nil, err
	}
}

// Delete project.
func (project Project) Delete() error {
	if db, err := database.Open(); err == nil {
		db.Model(&project).Association("Tags").Clear()
		db.Model(&project).Association("Assets").Clear()
		db.Delete(&project)
		return nil
	} else {
		return err
	}
}
