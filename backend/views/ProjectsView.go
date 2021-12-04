package views

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/kataras/iris/v12"
)

func Projects(ctx iris.Context) {
	err := ctx.URLParam("err")
	type ProjectData struct {
		ID          string
		Index       int
		Title       string
		Description string
		DateCreated string
		DateUpdated string
	}

	projects, projects_err := models.GetProjects()
	if projects_err != nil {
		err = projects_err.Error()
	}

	var project_data []*ProjectData
	layout := "2006-01-02 15:04:05 -0700 MST"
	zone, _ := time.LoadLocation("Australia/Perth")

	for i, project := range projects {
		DateCreated, _ := time.Parse(layout, project.DateCreated)
		DateUpdated, _ := time.Parse(layout, project.DateUpdated)

		project_data = append(project_data, &ProjectData{
			ID:          project.ID.Hex(),
			Index:       i + 1,
			Title:       project.Title,
			Description: project.Description,
			DateCreated: DateCreated.In(zone).Format(time.RFC822),
			DateUpdated: DateUpdated.In(zone).Format(time.RFC822),
		})
	}

	helpers.RenderTemplate(ctx, "projects/projects", "admin", iris.Map{"Projects": project_data, "Err": err})
}

func NewProject(ctx iris.Context) {
	if tags, tags_err := models.GetTags(); tags_err != nil {
		helpers.RenderTemplate(ctx, "projects/create", "admin", iris.Map{"Tags": []models.Tag{}, "Err": tags_err.Error()})
	} else {
		helpers.RenderTemplate(ctx, "projects/create", "admin", iris.Map{"Tags": tags})
	}
}

func EditProject(ctx iris.Context) {
	type TagData struct {
		Title   string
		Slug    string
		Checked bool
	}

	type ProjectData struct {
		ID          string
		Title       string
		Slug        string
		Description string
		Image       string
		URL         string
		Git         string
		Tags        []*models.Tag
		DateCreated string
		DateUpdated string
	}

	id := ctx.Params().Get("id")
	project, project_err := models.GetProject(id)

	if project_err != nil {
		ctx.View("404")
		return
	}

	project_url := ""
	if project.URL != nil {
		project_url = *project.URL
	}

	git_url := ""
	if project.Git != nil {
		git_url = *project.Git
	}

	project_data := ProjectData{
		ID:          project.ID.Hex(),
		Title:       project.Title,
		Slug:        project.Slug,
		Image:       project.Image,
		URL:         project_url,
		Git:         git_url,
		Tags:        project.Tags,
		Description: project.Description,
		DateCreated: project.DateCreated,
	}

	var tags_data []*TagData
	tags, _ := models.GetTags()
	for _, tag := range tags {
		tags_data = append(tags_data, &TagData{
			Title:   tag.Title,
			Slug:    tag.Slug,
			Checked: models.CheckTagExistsInTags(tag, project.Tags),
		})
	}

	helpers.RenderTemplate(ctx, "projects/view", "admin", iris.Map{"Project": project_data, "Tags": tags_data})
}
