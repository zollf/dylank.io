package config

import (
	"app/graphql"
	"app/helpers"
	"app/middleware"
	"app/migrations"
	"app/routes"
	"app/utils"
	"fmt"
	"os"

	"github.com/kataras/iris/v12"
)

func Main() *iris.Application {
	DotEnv()
	if os.Getenv("MIGRATE") == "Yes" {
		migrations.Install()
	}

	utils.Log().Info("Application starting...")

	root := helpers.RootDir()
	app := iris.New()

	app.RegisterView(iris.Blocks(fmt.Sprintf("%s/resources/templates", root), ".html").Reload(true))
	app.HandleDir("/admin/styles", iris.Dir(fmt.Sprintf("%s/resources/static/styles", root)))

	app.Post("/api/graphql", graphql.ExecuteGraphqlQuery)

	routes.AuthRoutes(app)

	app.Use(middleware.AuthRequired)

	routes.AdminRoutes(app)
	routes.UserRoutes(app)
	routes.ProjectsRoutes(app)
	routes.TagsRoutes(app)
	routes.AssetsRoutes(app)

	return app
}
