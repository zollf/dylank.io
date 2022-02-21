package views

import (
	"app/helpers"
	"app/models"

	"github.com/kataras/iris/v12"
)

func Assets(ctx iris.Context) {
	err := ctx.URLParam("err")
	assets, _ := models.GetAssetsData()

	helpers.RenderTemplate(ctx, "assets/assets", "admin", iris.Map{"Assets": assets, "Err": err})
}

func NewAsset(ctx iris.Context) {
	helpers.RenderTemplate(ctx, "assets/create", "admin", iris.Map{})
}

func EditAsset(ctx iris.Context) {
	id := ctx.Params().Get("id")
	asset, not_found := models.GetAssetData(id)

	if not_found != nil {
		ctx.View("404")
		return
	}

	helpers.RenderTemplate(ctx, "assets/view", "admin", iris.Map{"Asset": asset})
}
