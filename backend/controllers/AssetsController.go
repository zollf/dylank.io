package controllers

import (
	"app/helpers"
	"app/models"

	"github.com/gosimple/slug"
	"github.com/kataras/iris/v12"
)

func CreateAsset(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"title"}) {
		return
	}

	files, files_err := helpers.UploadImage(ctx, "image")
	if files_err != nil {
		helpers.ErrorResponse(ctx, "Failed to upload files to s3", iris.Map{"error": files_err.Error(), "files": files})
		return
	}

	if len(files) == 0 {
		helpers.ErrorResponse(ctx, "Internal Error", iris.Map{"files": files})
		return
	}

	file := files[0]

	asset := &models.Asset{
		Title: file.Title,
		Slug:  slug.Make(file.Title),
		Url:   file.Url,
	}

	if err := models.CreateAsset(asset); err != nil {
		helpers.ErrorResponse(ctx, "Failed to upload asset", iris.Map{"error": err.Error()})
	} else {
		helpers.SuccessResponse(ctx, "Successfully created project", iris.Map{"asset": asset})
	}
}

func DeleteAsset(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"id"}) {
		return
	}

	if err := models.DeleteAsset(ctx.FormValue("id")); err != nil {
		helpers.ErrorResponse(ctx, "Failed to delete asset", iris.Map{"error": err.Error()})
	} else {
		helpers.SuccessResponse(ctx, "Successfully deleted asset", iris.Map{})
	}
}
