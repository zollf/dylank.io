package controllers

import "github.com/kataras/iris/v12"

func Login(ctx iris.Context) {
	ctx.View("login.html", iris.Map{})
}
