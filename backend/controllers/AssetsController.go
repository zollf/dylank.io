package controllers

import (
	"app/helpers"
	"app/helpers/res"
	"app/models/assets"

	"github.com/gosimple/slug"
	"github.com/kataras/iris/v12"
)

type AssetsController struct{}

// Get Assets
// Method:   GET
// Resource: /api/assets
func (c *AssetsController) Get(ctx iris.Context) {
	if assets, cannot_list := assets.All(); cannot_list != nil {
		res.ASSETS_LIST.Error(ctx, cannot_list)
	} else {
		res.ASSETS_LIST.Send(ctx, iris.Map{"assets": assets})
	}
}

// Creates Assets
// Method:   POST
// Resource: /api/assets/create
func (c *AssetsController) PostCreate(req assets.AssetCreateRequest, ctx iris.Context) {
	res.ASSET_CREATE.Validate(ctx, &req)
	if ctx.IsStopped() {
		return
	}

	file, upload_err := helpers.UploadImage(ctx, "image", req.Title)
	if upload_err != nil {
		res.ASSET_CREATE.Error(ctx, upload_err)
		return
	}

	asset := &assets.Asset{
		Title: file.Title,
		Slug:  slug.Make(file.Title),
		Url:   file.Url,
	}

	if asset_create_err := asset.Create(); asset_create_err != nil {
		res.ASSET_CREATE.Error(ctx, asset_create_err)
	} else {
		res.ASSET_CREATE.Send(ctx, iris.Map{"asset": asset})
	}
}

// Delete Assets
// Method:   POST
// Resource: /api/assets/delete
func (c *AssetsController) PostDelete(req assets.AssetDeleteRequest, ctx iris.Context) {
	res.ASSET_DELETE.Validate(ctx, &req)
	if ctx.IsStopped() {
		return
	}

	asset, not_found := assets.Find(req.ID)
	if not_found != nil {
		res.ASSET_DELETE.Error(ctx, not_found)
		return
	}

	if delete_error := asset.Delete(); delete_error != nil {
		res.ASSET_DELETE.Error(ctx, delete_error)
	} else {
		res.ASSET_DELETE.Send(ctx, iris.Map{})
	}
}
