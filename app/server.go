package main

import (
	"app/config"
	"app/controllers"

	"github.com/kataras/iris/v12"
)

func main() {
	config.DotEnv()
	app := iris.New()
	app.WrapRouter(config.Graphql)

	e := iris.Django("./views", ".html").Reload(true)
	app.RegisterView(e)

	app.HandleDir("/admin/styles", iris.Dir("./views/styles"))

	app.Get("/admin", controllers.AdminIndex)
	app.Get("/admin/login", controllers.Login)

	app.Listen(":8080")
}
