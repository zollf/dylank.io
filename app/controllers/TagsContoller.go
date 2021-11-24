package controllers

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/gosimple/slug"

	"github.com/kataras/iris/v12"
)

func CreateOrEditTag(ctx iris.Context) {
	if ctx.FormValue("title") == "" {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Please include title and slug"), nil, iris.Map{"tag": nil})
		return
	}

	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to get or create id"), nil, iris.Map{"tag": nil})
	} else {
		tag := &models.Tag{
			ID:          id,
			Title:       ctx.FormValue("title"),
			Slug:        slug.Make(ctx.FormValue("title")),
			DateCreated: helpers.GetOrCreateDate(ctx),
			DateUpdated: time.Now().UTC().String(),
		}

		if err := models.CreateOrEditTag(tag); err != nil {
			helpers.RedirectIfExist(ctx, helpers.ErrorMsg(err.Error()), nil, iris.Map{"tag": nil})
		} else {
			helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully saved tag"), iris.Map{"tag": tag})
		}
	}
}

func DeleteTag(ctx iris.Context) {
	if ctx.FormValue("id") == "" {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Please include id of tag"), nil, iris.Map{})
		return
	}

	if err := models.DeleteTag(ctx.FormValue("id")); err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to delete tag"), nil, iris.Map{})
	} else {
		helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully deleted tag"), iris.Map{})
	}
}
