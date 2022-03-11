package controllers

import (
	"app/helpers/res"
	"app/models"
	"app/services"
	"errors"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	type Req struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	if !res.AUTH_LOGIN.Validate(ctx, &req) {
		return
	}

	if user, err := models.GetUserWithPassword(req.Username, req.Password); err != nil {
		res.AUTH_LOGIN.Error(ctx, errors.New("Username or password is incorrect"))
	} else {
		// Generate JWT Token
		if token, err := services.GenerateJWT(user); err != nil {
			res.AUTH_LOGIN.Error(ctx, errors.New("Failed to generate token"))
		} else {
			ctx.SetCookie(&iris.Cookie{
				Name:   "dylank-io-auth",
				Value:  token,
				Secure: true,
			}, iris.CookiePath("/"))
			res.AUTH_LOGIN.Send(ctx, iris.Map{"token": token})
		}
	}
}

func Logout(ctx iris.Context) {
	ctx.RemoveCookie("dylank-io-auth")
	res.AUTH_LOGOUT.Send(ctx, iris.Map{})
}
