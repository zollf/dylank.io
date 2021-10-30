package main

import (
	"app/config"
	"app/controllers"
	"app/middleware"
	"app/views"

	"github.com/kataras/iris/v12"
)

func main() {
	config.DotEnv()
	app := iris.New()
	app.WrapRouter(config.Graphql)

	e := iris.Pug("./resources/templates", ".pug").Reload(true)
	app.RegisterView(e)
	app.HandleDir("/admin/styles", iris.Dir("./resources/static/styles"))

	app.Get("/admin/login", views.Login)
	app.Get("/admin/logout", views.Logout)
	app.Get("/api/logout", controllers.Logout)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.AuthRequired)
	app.Get("/admin", views.AdminIndex)

	app.Get("/admin/users/create", views.NewUser)
	app.Get("/admin/users/view/{id}", views.EditUser)
	app.Get("/admin/users", views.Users)

	app.Post("/api/users/create", controllers.CreateOrEditUser)
	app.Post("/api/users/edit", controllers.CreateOrEditUser)
	app.Post("/api/users/delete", controllers.DeleteUser)

	app.Get("/admin/projects", views.Projects)
	app.Get("/admin/projects/create", views.NewProject)

	app.Listen(":8080")
}
