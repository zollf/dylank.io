package res

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func Redirect(ctx iris.Context, params string) bool {
	if redirect := ctx.FormValue("redirect"); redirect != "" {
		ctx.Redirect(fmt.Sprintf("%s%s", redirect, params))
		return true
	}
	return false
}

func AcceptsJSON(ctx iris.Context) bool {
	return ctx.GetHeader("Accepts") == "application/json"
}
