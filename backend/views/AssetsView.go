package views

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/kataras/iris/v12"
)

func Assets(ctx iris.Context) {
	err := ctx.URLParam("err")
	type AssetData struct {
		ID        uint64
		Index     int
		Title     string
		Url       string
		CreatedAt string
		UpdatedAt string
	}

	assets, assets_err := models.GetAssets()
	if assets_err != nil {
		err = assets_err.Error()
	}

	var asset_data []*AssetData
	zone, _ := time.LoadLocation("Australia/Perth")

	for i, asset := range assets {
		asset_data = append(asset_data, &AssetData{
			ID:        asset.ID,
			Index:     i + 1,
			Url:       asset.Url,
			Title:     asset.Title,
			CreatedAt: asset.CreatedAt.In(zone).Format(time.RFC822),
			UpdatedAt: asset.UpdatedAt.In(zone).Format(time.RFC822),
		})
	}

	helpers.RenderTemplate(ctx, "assets/assets", "admin", iris.Map{"Assets": asset_data, "Err": err})
}

func NewAsset(ctx iris.Context) {
	helpers.RenderTemplate(ctx, "assets/create", "admin", iris.Map{})
}

func EditAsset(ctx iris.Context) {
	type AssetData struct {
		ID        uint64
		Title     string
		Slug      string
		Url       string
		CreatedAt string
		UpdatedAt string
	}

	id := ctx.Params().Get("id")
	asset, asset_err := models.GetAsset(id)

	if asset_err != nil {
		ctx.View("404")
		return
	}

	zone, _ := time.LoadLocation("Australia/Perth")

	asset_data := &AssetData{
		ID:        asset.ID,
		Title:     asset.Title,
		Url:       asset.Url,
		CreatedAt: asset.CreatedAt.In(zone).Format(time.RFC822),
		UpdatedAt: asset.UpdatedAt.In(zone).Format(time.RFC822),
	}

	helpers.RenderTemplate(ctx, "assets/view", "admin", iris.Map{"Asset": asset_data})
}
