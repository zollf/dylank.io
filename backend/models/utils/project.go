package utils

import (
	"app/models/assets"
	"app/models/projects"
	"app/models/tags"
	"time"
)

// Used for dom rendering
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
	Assets      []*assets.Asset
	URL         string
	Git         string
	Tags        []*tags.Tag
	CreatedAt   string
	UpdatedAt   string
}

func GetProjectData() ([]*ProjectData, error) {
	projects, projects_err := projects.All()
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
	project, project_err := projects.Find(id)

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
	tags, _ := tags.All()
	for _, tag := range tags {
		tags_data = append(tags_data, &ProjectTagData{
			Title:   tag.Title,
			Slug:    tag.Slug,
			Checked: CheckTagExistsInTags(tag, project.Tags),
		})
	}

	var assets_data []*ProjectAssetData
	assets, _ := assets.All()
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
