package main

import (
	"app/config"
	"app/middleware"
	"app/routes"
	"app/scripts"
	"os"

	"github.com/kataras/iris/v12"
)

func main() {
	config.DotEnv()

	// Run server if there is no other args
	if len(os.Args) == 1 || os.Args[1] == "runserver" {
		app := iris.New()
		app.WrapRouter(config.Graphql)

		e := iris.Pug("./resources/templates", ".pug").Reload(true)
		app.RegisterView(e)
		app.HandleDir("/admin/styles", iris.Dir("./resources/static/styles"))

		routes.AuthRoutes(app)

		app.Use(middleware.AuthRequired)

		routes.AdminRoutes(app)
		routes.UserRoutes(app)
		routes.ProjectsRoutes(app)

		app.Listen(":8080")
	} else {
		scripts.RunScripts(os.Args[1:])
	}
}
