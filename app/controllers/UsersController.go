package controllers

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/kataras/iris/v12"
)

func CreateOrEditUser(ctx iris.Context) {
	if invalidUserRequest(ctx) {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Params invalid"), nil, iris.Map{})
	}

	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to get or create id"), nil, iris.Map{})
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
			helpers.RedirectIfExist(ctx, helpers.ErrorMsg(err.Error()), nil, iris.Map{})

		} else {
			helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully saved user"), iris.Map{})
		}
	}
}

func DeleteUser(ctx iris.Context) {
	if ctx.FormValue("id") == "" {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Please include id of user"), nil, iris.Map{})
		return
	}

	if err := models.DeleteUser(ctx.FormValue("id")); err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to delete user"), nil, iris.Map{})
	} else {
		helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully deleted user"), iris.Map{})
	}
}

func invalidUserRequest(ctx iris.Context) bool {
	return ctx.FormValue("username") == "" || ctx.FormValue("password") == "" || ctx.FormValue("email") == ""
}
