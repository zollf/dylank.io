package routes

import (
	"app/views"

	"github.com/kataras/iris/v12"
)

func ProjectsRoutes(app *iris.Application) {
	app.Get("/admin/projects", views.Projects)
	app.Get("/admin/projects/create", views.NewProject)
}
