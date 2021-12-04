package views

import (
	"app/helpers"
	"app/services"

	"github.com/kataras/iris/v12"
)

func AdminIndex(ctx iris.Context) {
	username, _ := services.GetAndVerifyCookie(ctx)
	helpers.RenderTemplate(ctx, "index", "admin", iris.Map{"Username": username})
}
