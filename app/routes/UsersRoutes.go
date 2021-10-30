package routes

import (
	"app/controllers"
	"app/views"

	"github.com/kataras/iris/v12"
)

func UserRoutes(app *iris.Application) {
	app.Get("/admin/users/create", views.NewUser)
	app.Get("/admin/users/view/{id}", views.EditUser)
	app.Get("/admin/users", views.Users)

	app.Post("/api/users/create", controllers.CreateOrEditUser)
	app.Post("/api/users/edit", controllers.CreateOrEditUser)
	app.Post("/api/users/delete", controllers.DeleteUser)
}
