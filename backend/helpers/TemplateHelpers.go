package helpers

import (
	"strings"

	"github.com/kataras/iris/v12"
)

func RenderTemplate(ctx iris.Context, template string, layout string, data iris.Map) {
	if data["Err"] == nil {
		err := ctx.URLParam("err")
		data["Err"] = err
	}

	if data["Success"] == nil {
		success := ctx.URLParam("success")
		data["Success"] = success
	}

	data["Path"] = ctx.Path()

	pathSliced := strings.Split(ctx.Path(), "/")
	data["EndPath"] = pathSliced[len(pathSliced)-1]

	if token := GetToken(ctx); token != "" {
		data["LoggedIn"] = true
	} else {
		data["LoggedIn"] = false
	}

	ctx.ViewLayout(layout)
	ctx.View(template, data)
}
