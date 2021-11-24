package controllers

import (
	"app/helpers"
	"app/models"
	"app/services"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	var err_msg *string
	var success_msg *string
	var token_var *string

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if user, err := models.GetUserWithPassword(username, password); err != nil {
		err_msg = helpers.ErrorMsg("Incorrect credentials")
	} else {
		// Generate JWT Token
		if token, err := services.GenerateJWT(user); err != nil {
			err_msg = helpers.ErrorMsg("Failed to generate token")
		} else {
			ctx.SetCookie(&iris.Cookie{
				Name:   "dylank-io-auth",
				Value:  token,
				Secure: true,
			}, iris.CookiePath("/"))

			success_msg = helpers.SuccessMsg("Successfully created token")
			token_var = &token
		}
	}

	helpers.RedirectIfExist(ctx, err_msg, success_msg, iris.Map{
		"token": token_var,
	})
}

func Logout(ctx iris.Context) {
	ctx.RemoveCookie("dylank-io-auth")
	ctx.Redirect("/admin/login?success=Logout+Successful")
}
