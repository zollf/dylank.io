package controllers

import (
	"app/helpers"
	"app/models"
	"app/services"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"username", "password"}) {
		return
	}

	if user, err := models.GetUserWithPassword(ctx.FormValue("username"), ctx.FormValue("password")); err != nil {
		helpers.ErrorResponse(ctx, "Incorrect credentials", iris.Map{"error": err.Error()})
	} else {
		// Generate JWT Token
		if token, err := services.GenerateJWT(user); err != nil {
			helpers.ErrorResponse(ctx, "Failed to generate token", iris.Map{"error": err.Error()})
		} else {
			ctx.SetCookie(&iris.Cookie{
				Name:   "dylank-io-auth",
				Value:  token,
				Secure: true,
			}, iris.CookiePath("/"))

			helpers.SuccessResponse(ctx, "Successfully created token", iris.Map{"token": token})
		}
	}
}

func Logout(ctx iris.Context) {
	ctx.RemoveCookie("dylank-io-auth")
	ctx.Redirect("/admin/login?success=Logout+Successful")
}
