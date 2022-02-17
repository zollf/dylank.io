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
		ID          uint64
		Index       int
		Title       string
		Description string
		CreatedAt   string
		UpdatedAt   string
	}

	projects, projects_err := models.GetProjects()
	if projects_err != nil {
		err = projects_err.Error()
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

	helpers.RenderTemplate(ctx, "projects/projects", "admin", iris.Map{"Projects": project_data, "Err": err})
}

func NewProject(ctx iris.Context) {
	type TagData struct {
		Title string
		Slug  string
	}

	var tags_data []*TagData
	tags, _ := models.GetTags()
	for _, tag := range tags {
		tags_data = append(tags_data, &TagData{
			Title: tag.Title,
			Slug:  tag.Slug,
		})
	}

	helpers.RenderTemplate(ctx, "projects/create", "admin", iris.Map{"Tags": tags_data})
}

func EditProject(ctx iris.Context) {
	type TagData struct {
		Title   string
		Slug    string
		Checked bool
	}

	type ProjectData struct {
		ID          uint64
		Title       string
		Slug        string
		Description string
		Image       string
		URL         string
		Git         string
		Tags        []*models.Tag
		CreatedAt   string
		UpdatedAt   string
	}

	id := ctx.Params().Get("id")
	project, project_err := models.GetProject(id)

	if project_err != nil {
		ctx.View("404")
		return
	}

	zone, _ := time.LoadLocation("Australia/Perth")

	project_data := &ProjectData{
		ID:          project.ID,
		Title:       project.Title,
		Slug:        project.Slug,
		Image:       project.Image,
		URL:         helpers.StringLike(project.URL),
		Git:         helpers.StringLike(project.Git),
		Tags:        project.Tags,
		Description: project.Description,
		CreatedAt:   project.CreatedAt.In(zone).Format(time.RFC822),
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
