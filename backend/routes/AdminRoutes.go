package routes

import (
	"app/views"

	"github.com/kataras/iris/v12"
)

func AdminRoutes(app *iris.Application) {
	app.Get("/admin", views.AdminIndex)
}
