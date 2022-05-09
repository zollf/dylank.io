package routes

import (
	"app/views"

	"github.com/kataras/iris/v12"
)

func TagsRoutes(app *iris.Application) {
	app.Get("/admin/tags", views.Tags)
	app.Get("/admin/tags/create", views.NewTag)
	app.Get("/admin/tags/view/{id}", views.EditTag)
}
