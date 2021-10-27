package controllers

import (
	"app/models"
	"app/services"
	"fmt"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	redirect := ctx.FormValue("redirect")

	user, err := models.GetUserWithPassword(username, password)

	if err != nil {
		ctx.Redirect("/admin/login?err=Incorrect credentials")
		return
	} else {
		// Generate JWT Token
		token, err := services.GenerateJWT(user)

		if err != nil {
			if redirect != "" {
				ctx.Redirect(fmt.Sprintf("/admin/login?err=Error occurred&redirect=%s", redirect))
				return
			} else {
				ctx.Redirect("/admin/login?err=Error occurred")
				return
			}
		}

		ctx.SetCookie(&iris.Cookie{
			Name:   "dylank-io-auth",
			Value:  token,
			Secure: true,
		}, iris.CookiePath("/"))

		if redirect != "" {
			ctx.Redirect(redirect)
			return
		}
		ctx.Redirect("/admin?success=true")
		return
	}
}

func Logout(ctx iris.Context) {
	ctx.RemoveCookie("dylank-io-auth")
	ctx.Redirect("/admin/login?success=Logout Successful")
}
