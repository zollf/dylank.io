package controllers

import (
	"app/helpers"
	"app/models"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

func CreateOrEditUser(ctx iris.Context) {
	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/projects?err=%s", id_err.Error()))
	} else {
		user := &models.User{
			ID:           id,
			Username:     ctx.FormValue("username"),
			Password:     ctx.FormValue("password"),
			Email:        ctx.FormValue("email"),
			Locked:       false,
			DateCreated:  helpers.GetOrCreateDate(ctx),
			DateUpdated:  time.Now().UTC().String(),
			LastLoggedIn: "",
		}
		if err := models.CreateOrEditUser(user); err != nil {
			ctx.Redirect(fmt.Sprintf("/admin/users?err=%s", err.Error()))
		} else {
			ctx.Redirect("/admin/users?success=true")
		}
	}
}

func DeleteUser(ctx iris.Context) {
	if err := models.DeleteUser(ctx.FormValue("id")); err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/users?err=%s", err.Error()))
	} else {
		ctx.Redirect("/admin/users?success=true")
	}
}
