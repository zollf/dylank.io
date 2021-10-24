package controllers

import "github.com/kataras/iris/v12"

func AdminIndex(ctx iris.Context) {
	ctx.View("index.html", iris.Map{
		"title": "Hi Page",
		"name":  "iris",
	})
}
