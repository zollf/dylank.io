package controllers

import (
	"app/helpers"
	"app/helpers/res"
	"app/models/assets"

	"github.com/gosimple/slug"
	"github.com/kataras/iris/v12"
)

func AssetsList(ctx iris.Context) {
	if assets, cannot_list := assets.All(); cannot_list != nil {
		res.ASSETS_LIST.Error(ctx, cannot_list)
	} else {
		res.ASSETS_LIST.Send(ctx, iris.Map{"assets": assets})
	}
}

func AssetsCreate(ctx iris.Context) {
	type Req struct {
		Title    string `json:"title" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	if !res.ASSETS_CREATE.Validate(ctx, &req) {
		return
	}

	file, upload_err := helpers.UploadImage(ctx, "image", req.Title)
	if upload_err != nil {
		res.ASSETS_CREATE.Error(ctx, upload_err)
		return
	}

	asset := &assets.Asset{
		Title: file.Title,
		Slug:  slug.Make(file.Title),
		Url:   file.Url,
	}

	if asset_create_err := asset.Create(); asset_create_err != nil {
		res.ASSETS_CREATE.Error(ctx, asset_create_err)
	} else {
		res.ASSETS_CREATE.Send(ctx, iris.Map{"asset": asset})
	}
}

func AssetsDelete(ctx iris.Context) {
	type Req struct {
		ID       string `json:"id" validate:"required"`
		Title    string `json:"title" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	if !res.ASSETS_DELETE.Validate(ctx, &req) {
		return
	}

	asset, not_found := assets.Find(req.ID)
	if not_found != nil {
		res.ASSETS_DELETE.Error(ctx, not_found)
		return
	}

	if delete_error := asset.Delete(); delete_error != nil {
		res.ASSETS_DELETE.Error(ctx, delete_error)
	} else {
		res.ASSETS_DELETE.Send(ctx, iris.Map{})
	}
}
