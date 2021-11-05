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
		ctx.Redirect("/admin/login?err=Incorrect credentials")
	} else {
		// Generate JWT Token
		if token, err := services.GenerateJWT(user); err != nil {
			helpers.SaveRedirectIfExist(ctx, "/admin/login?err=Error occurred", "&")
		} else {
			ctx.SetCookie(&iris.Cookie{
				Name:   "dylank-io-auth",
				Value:  token,
				Secure: true,
			}, iris.CookiePath("/"))

			helpers.RedirectIfExist(ctx, "/admin?success=true")
		}
	}
}

func Logout(ctx iris.Context) {
	ctx.RemoveCookie("dylank-io-auth")
	ctx.Redirect("/admin/login?success=Logout Successful")
}
