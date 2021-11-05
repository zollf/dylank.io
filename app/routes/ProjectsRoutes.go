package routes

import (
	"app/controllers"
	"app/views"

	"github.com/kataras/iris/v12"
)

func ProjectsRoutes(app *iris.Application) {
	app.Get("/admin/projects", views.Projects)
	app.Get("/admin/projects/create", views.NewProject)
	app.Get("/admin/projects/view/{id}", views.EditProject)

	app.Post("/api/projects/create", controllers.CreateOrEditProject)
	app.Post("/api/projects/edit", controllers.CreateOrEditProject)
	app.Post("/api/projects/delete", controllers.DeleteProject)
}
