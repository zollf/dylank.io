package controllers

import (
	"app/helpers"
	"app/models"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

func CreateOrEditTag(ctx iris.Context) {
	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/tags?err=%s", id_err.Error()))
	} else {

		tag := &models.Tag{
			ID:          id,
			Title:       ctx.FormValue("title"),
			Slug:        ctx.FormValue("slug"),
			DateCreated: helpers.GetOrCreateDate(ctx),
			DateUpdated: time.Now().UTC().String(),
		}

		if err := models.CreateOrEditTag(tag); err != nil {
			ctx.Redirect(fmt.Sprintf("/admin/tags?err=%s", err.Error()))
		} else {
			ctx.Redirect("/admin/tags?success=true")
		}
	}
}

func DeleteTag(ctx iris.Context) {
	if err := models.DeleteTag(ctx.FormValue("id")); err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/tags?err=%s", err.Error()))
	} else {
		ctx.Redirect("/admin/tags?success=true")
	}
}
