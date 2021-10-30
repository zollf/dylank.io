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

	projects, _ := models.GetProjects()
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

func EditProject() {

}
