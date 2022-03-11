package routes

import (
	"app/controllers"
	"app/views"

	"github.com/kataras/iris/v12"
)

func AuthRoutes(app *iris.Application) {
	app.Get("/admin/login", views.Login)
	app.Get("/admin/logout", views.Logout)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/login", controllers.Login)
}
