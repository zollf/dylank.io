package controllers

import (
	"app/helpers"
	"app/models"

	"github.com/gosimple/slug"

	"github.com/kataras/iris/v12"
)

func CreateTag(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"title"}) {
		return
	}

	tag := &models.Tag{
		Title: ctx.FormValue("title"),
		Slug:  slug.Make(ctx.FormValue("title")),
	}

	if err := models.CreateTag(tag); err != nil {
		helpers.ErrorResponse(ctx, "Failed to created tag", iris.Map{"error": err.Error()})
	} else {
		helpers.SuccessResponse(ctx, "Successfully created tag", iris.Map{"tag": tag})
	}
}

func EditTag(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"id", "title"}) {
		return
	}

	tag := &models.Tag{
		Title: ctx.FormValue("title"),
		Slug:  slug.Make(ctx.FormValue("title")),
	}

	if err := models.UpdateTag(tag, ctx.FormValue("id")); err != nil {
		helpers.ErrorResponse(ctx, "Failed to update tag", iris.Map{"error": err.Error()})
	} else {
		helpers.SuccessResponse(ctx, "Successfully updated tag", iris.Map{"tag": tag})
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
