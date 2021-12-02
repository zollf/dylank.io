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
			user.Password = "***"
			helpers.SuccessResponse(ctx, "Successfully saved user", iris.Map{"user": user})
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

func ListUsers(ctx iris.Context) {
	users, err := models.GetUsers()
	if err != nil {
		helpers.ErrorResponse(ctx, "Failed to list users", iris.Map{"error": err.Error()})
	}
	for _, user := range users {
		user.Password = "***"
	}
	helpers.SuccessResponse(ctx, "Successfully listed users", iris.Map{"users": users})
}
