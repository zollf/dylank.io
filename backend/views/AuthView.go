package views

import (
	"app/services"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	err := ctx.URLParam("err")
	success := ctx.URLParam("success")

	_, verify_err := services.GetAndVerifyCookie(ctx)

	if verify_err != nil {
		ctx.View("auth/login.pug", iris.Map{"Err": err, "Success": success})
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
		ctx.View("auth/logout.pug", iris.Map{"Username": username})
		return
	}
}
