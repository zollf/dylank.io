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
	Assets      []*Asset  `json:"assets" gorm:"many2many:project_assets"`
	URL         *string   `json:"url"`
	Git         *string   `json:"git"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Tags        []*Tag    `json:"tags" gorm:"many2many:project_tags"`
}

type ProjectData struct {
	ID          uint64
	Index       int
	Title       string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type ProjectDataWithTagAndAsset struct {
	ID          uint64
	Title       string
	Slug        string
	Description string
	Assets      []*Asset
	URL         string
	Git         string
	Tags        []*Tag
	CreatedAt   string
	UpdatedAt   string
}

func GetProjects() ([]*Project, error) {
	var projects []*Project
	if db, err := database.Open(); err == nil {
		results := db.Preload("Tags").Preload("Assets").Find(&projects)
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

func CreateProject(project *Project) error {
	return database.CreateRecord(project)
}

func GetProject(id string) (*Project, error) {
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

func DeleteProject(id string) error {
	return database.DeleteRecord(&Project{}, id)
}

func GetProjectData() ([]*ProjectData, error) {
	projects, projects_err := GetProjects()
	if projects_err != nil {
		return nil, projects_err
	}

	var project_data []*ProjectData
	zone, _ := time.LoadLocation("Australia/Perth")

	for i, project := range projects {
		project_data = append(project_data, &ProjectData{
			ID:          project.ID,
			Index:       i + 1,
			Title:       project.Title,
			Description: project.Description,
			CreatedAt:   project.CreatedAt.In(zone).Format(time.RFC822),
			UpdatedAt:   project.UpdatedAt.In(zone).Format(time.RFC822),
		})
	}

	return project_data, nil
}

func GetProjectDataWithTagAndAsset(id string) (*ProjectDataWithTagAndAsset, []*ProjectTagData, []*ProjectAssetData, error) {
	project, project_err := GetProject(id)

	if project_err != nil {
		return nil, nil, nil, project_err
	}

	zone, _ := time.LoadLocation("Australia/Perth")

	Git := ""
	if project.Git != nil {
		Git = *project.Git
	}

	URL := ""
	if project.URL != nil {
		URL = *project.URL
	}

	project_data := &ProjectDataWithTagAndAsset{
		ID:          project.ID,
		Title:       project.Title,
		Slug:        project.Slug,
		Assets:      project.Assets,
		URL:         URL,
		Git:         Git,
		Tags:        project.Tags,
		Description: project.Description,
		CreatedAt:   project.CreatedAt.In(zone).Format(time.RFC822),
	}

	var tags_data []*ProjectTagData
	tags, _ := GetTags()
	for _, tag := range tags {
		tags_data = append(tags_data, &ProjectTagData{
			Title:   tag.Title,
			Slug:    tag.Slug,
			Checked: CheckTagExistsInTags(tag, project.Tags),
		})
	}

	var assets_data []*ProjectAssetData
	assets, _ := GetAssets()
	for _, asset := range assets {
		assets_data = append(assets_data, &ProjectAssetData{
			Title:   asset.Title,
			Slug:    asset.Slug,
			Url:     asset.Url,
			Checked: CheckAssetExistsInAssets(asset, project.Assets),
		})
	}

	return project_data, tags_data, assets_data, nil
}
