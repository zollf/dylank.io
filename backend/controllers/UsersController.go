package controllers

import (
	"app/helpers"
	"app/models"

	"github.com/kataras/iris/v12"
)

func CreateUser(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"username", "password", "email"}) {
		return
	}

	user := &models.User{
		Username:     ctx.FormValue("username"),
		Password:     ctx.FormValue("password"),
		Email:        ctx.FormValue("email"),
		Locked:       false,
		LastLoggedIn: nil,
	}

	if err := models.CreateUser(user); err != nil {
		helpers.ErrorResponse(ctx, "Failed to created user", iris.Map{"error": err.Error()})
	} else {
		user.Password = "***"
		helpers.SuccessResponse(ctx, "Successfully created user", iris.Map{"user": user})
	}
}

func EditUser(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"username", "password", "email", "id"}) {
		return
	}

	user := &models.User{
		Username:     ctx.FormValue("username"),
		Password:     ctx.FormValue("password"),
		Email:        ctx.FormValue("email"),
		Locked:       false,
		LastLoggedIn: nil,
	}

	if err := models.UpdateUser(user, ctx.FormValue("id")); err != nil {
		helpers.ErrorResponse(ctx, "Failed to update user", iris.Map{"error": err.Error()})
	} else {
		user.Password = "***"
		helpers.SuccessResponse(ctx, "Successfully updated user", iris.Map{"user": user})
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
