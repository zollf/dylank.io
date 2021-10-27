package views

import (
	"app/services"

	"github.com/kataras/iris/v12"
)

func AdminIndex(ctx iris.Context) {
	username, _ := services.GetAndVerifyCookie(ctx)
	ctx.View("index.pug", iris.Map{"Username": username})
}
