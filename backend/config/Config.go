package config

import (
	"app/controllers"
	"app/graphql"
	"app/middleware"
	"app/migrations"
	"app/routes"
	"app/utils"
	"fmt"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

const (
	PROD Env = "production"
	DEV  Env = "development"
	TEST Env = "test"
)

type Env string

func Main() *iris.Application {
	if os.Getenv("MIGRATE") == "Yes" {
		migrations.Install()
	}

	root := utils.RootDir()
	app := iris.New()

	app.RegisterView(iris.Blocks(fmt.Sprintf("%s/resources/templates", root), ".html").Reload(true))
	app.HandleDir("/admin/styles", iris.Dir(fmt.Sprintf("%s/resources/static/styles", root)))

	app.Post("/api/graphql", graphql.ExecuteGraphqlQuery)

	routes.AuthRoutes(app)

	app.Use(middleware.AuthRequired)

	mvc.New(app.Party("/api/tags")).Handle(new(controllers.TagsController))
	mvc.New(app.Party("/api/projects")).Handle(new(controllers.ProjectsController))
	mvc.New(app.Party("/api/assets")).Handle(new(controllers.AssetsController))
	mvc.New(app.Party("/api/users")).Handle(new(controllers.UsersController))

	routes.AdminRoutes(app)
	routes.UserRoutes(app)
	routes.ProjectsRoutes(app)
	routes.TagsRoutes(app)
	routes.AssetsRoutes(app)

	return app
}
