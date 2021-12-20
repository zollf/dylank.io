package helpers

import (
	"github.com/kataras/iris/v12"
)

func GetToken(ctx iris.Context) string {
	token := ctx.GetCookie("dylank-io-auth")
	if token == "" {
		token = ctx.FormValue("token")
	}

	if token == "" {
		token = ctx.GetHeader("Authentication")
	}

	return token
}
