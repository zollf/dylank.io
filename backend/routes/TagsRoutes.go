package routes

import (
	"app/controllers"
	"app/views"

	"github.com/kataras/iris/v12"
)

func TagsRoutes(app *iris.Application) {
	app.Get("/admin/tags", views.Tags)
	app.Get("/admin/tags/create", views.NewTag)
	app.Get("/admin/tags/view/{id}", views.EditTag)

	app.Get("/api/tags", controllers.ListTags)
	app.Post("/api/tags/create", controllers.CreateTag)
	app.Post("/api/tags/edit", controllers.EditTag)
	app.Post("/api/tags/delete", controllers.DeleteTag)
}
