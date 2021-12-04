package views

import (
	"app/helpers"
	"app/services"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	_, verify_err := services.GetAndVerifyCookie(ctx)

	if verify_err != nil {
		helpers.RenderTemplate(ctx, "auth/login", "base", iris.Map{})
		return
	} else {
		ctx.Redirect("/admin/logout")
		return
	}
}

func Logout(ctx iris.Context) {
	username, verify_err := services.GetAndVerifyCookie(ctx)

	if verify_err != nil {
		ctx.Redirect("/admin/login")
		return
	} else {
		helpers.RenderTemplate(ctx, "auth/logout", "base", iris.Map{"Username": username})
		return
	}
}
