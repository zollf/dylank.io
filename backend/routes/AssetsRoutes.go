package routes

import (
	"app/views"

	"github.com/kataras/iris/v12"
)

func AssetsRoutes(app *iris.Application) {
	app.Get("/admin/assets", views.Assets)
	app.Get("/admin/assets/create", views.NewAsset)
	app.Get("/admin/assets/view/{id}", views.EditAsset)
}
