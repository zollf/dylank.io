package views

import (
	"app/models"
	"time"

	"github.com/kataras/iris/v12"
)

func Projects(ctx iris.Context) {
	err := ctx.URLParam("err")
	success := ctx.URLParam("success")
	type ProjectData struct {
		ID          string
		Index       int
		Name        string
		Description string
		DateCreated string
		DateUpdated string
	}

	projects, p_err := models.GetProjects()
	if p_err != nil {
		err = p_err.Error()
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
			Name:        project.Title,
			Description: project.Description,
			DateCreated: DateCreated.In(zone).Format(time.RFC822),
			DateUpdated: DateUpdated.In(zone).Format(time.RFC822),
		})
	}

	ctx.View("projects/projects.pug", iris.Map{"Err": err, "Success": success, "Projects": project_data})
}

func NewProject(ctx iris.Context) {
	ctx.View("projects/create.pug")
}

func EditProject(ctx iris.Context) {
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
		ctx.View("404.pug")
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

	ctx.View("projects/view.pug", iris.Map{"Project": project_data})
}
