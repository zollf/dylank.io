package controllers

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/kataras/iris/v12"
)

func CreateOrEditUser(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"username", "password", "email"}) {
		return
	}

	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		helpers.ErrorResponse(ctx, "Failed to get or create id", iris.Map{"error": id_err.Error()})
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
			helpers.ErrorResponse(ctx, "Failed to save user", iris.Map{"error": err.Error()})
		} else {
			helpers.SuccessResponse(ctx, "Successfully saved user", iris.Map{})
		}
	}
}

func DeleteUser(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"id"}) {
		return
	}

	if err := models.DeleteUser(ctx.FormValue("id")); err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to delete user"), nil, iris.Map{"error": err.Error()})
	} else {
		helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully deleted user"), iris.Map{})
	}
}
