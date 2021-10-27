package controllers

import (
	"app/models"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrEditUser(ctx iris.Context) {
	var ID primitive.ObjectID
	if ctx.FormValue("id") != "" {
		pid, pid_err := primitive.ObjectIDFromHex(ctx.FormValue("id"))

		if pid_err != nil {
			ctx.Redirect(fmt.Sprintf("/admin/users?err=%s", pid_err.Error()))
			return
		}

		ID = pid
	} else {
		ID = primitive.NewObjectID()
	}

	var DateCreated string
	if ctx.FormValue("dateCreated") != "" {
		DateCreated = ctx.FormValue("dateCreated")
	} else {
		DateCreated = time.Now().UTC().String()
	}

	user := &models.User{
		ID:           ID,
		Username:     ctx.FormValue("username"),
		Password:     ctx.FormValue("password"),
		Email:        ctx.FormValue("email"),
		Locked:       false,
		DateCreated:  DateCreated,
		DateUpdated:  time.Now().UTC().String(),
		LastLoggedIn: "",
	}
	err := models.CreateOrEditUser(user)

	if err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/users?err=%s", err.Error()))
		return
	} else {
		ctx.Redirect("/admin/users?success=true")
		return
	}
}

func DeleteUser(ctx iris.Context) {
	err := models.DeleteUser(ctx.FormValue("id"))

	if err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/users?err=%s", err.Error()))
		return
	} else {
		ctx.Redirect("/admin/users?success=true")
		return
	}
}
