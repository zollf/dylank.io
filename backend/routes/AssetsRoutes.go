package routes

import (
	"app/controllers"
	"app/views"

	"github.com/kataras/iris/v12"
)

func AssetsRoutes(app *iris.Application) {
	app.Get("/admin/assets", views.Assets)
	app.Get("/admin/assets/create", views.NewAsset)
	app.Get("/admin/assets/view/{id}", views.EditAsset)

	app.Post("/api/assets/create", controllers.AssetsCreate)
	app.Post("/api/assets/delete", controllers.AssetsDelete)
	app.Get("/api/assets", controllers.AssetsList)
}
