package controllers

import (
	"app/helpers"
	"app/models"
	"app/services"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if user, err := models.GetUserWithPassword(username, password); err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Incorrect credentials"), nil, iris.Map{})
	} else {
		// Generate JWT Token
		if token, err := services.GenerateJWT(user); err != nil {
			helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to generate token"), nil, iris.Map{})
		} else {
			ctx.SetCookie(&iris.Cookie{
				Name:   "dylank-io-auth",
				Value:  token,
				Secure: true,
			}, iris.CookiePath("/"))

			helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully created token"), iris.Map{"token": token})
		}
	}
}

func Logout(ctx iris.Context) {
	ctx.RemoveCookie("dylank-io-auth")
	ctx.Redirect("/admin/login?success=Logout+Successful")
}
