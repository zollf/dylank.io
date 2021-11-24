package controllers

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/gosimple/slug"

	"github.com/kataras/iris/v12"
)

func CreateOrEditTag(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"title"}) {
		return
	}

	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		helpers.ErrorResponse(ctx, "Failed to get or create id", iris.Map{"tag": nil, "error": id_err.Error()})
	} else {
		tag := &models.Tag{
			ID:          id,
			Title:       ctx.FormValue("title"),
			Slug:        slug.Make(ctx.FormValue("title")),
			DateCreated: helpers.GetOrCreateDate(ctx),
			DateUpdated: time.Now().UTC().String(),
		}

		if err := models.CreateOrEditTag(tag); err != nil {
			helpers.ErrorResponse(ctx, "Failed to save tag", iris.Map{"tag": nil, "error": err.Error()})
		} else {
			helpers.SuccessResponse(ctx, "Successfully saved tag", iris.Map{"tag": tag})
		}
	}
}

func DeleteTag(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"id"}) {
		return
	}

	if err := models.DeleteTag(ctx.FormValue("id")); err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to delete tag"), nil, iris.Map{"error": err.Error()})
	} else {
		helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully deleted tag"), iris.Map{})
	}
}

func ListTags(ctx iris.Context) {
	tags, err := models.GetTags()
	if err != nil {
		helpers.ErrorResponse(ctx, "Failed to list tags", iris.Map{"error": err.Error()})
	}
	helpers.SuccessResponse(ctx, "Successfully listed tags", iris.Map{"tags": tags})
}
