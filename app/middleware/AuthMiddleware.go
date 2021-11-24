package middleware

import (
	"app/helpers"
	"app/services"
	"strings"

	"github.com/kataras/iris/v12"
)

// Auth is either through token param or cookie
func AuthRequired(ctx iris.Context) {
	if isApiRequest(ctx) {
		AuthApiRequest(ctx)
	} else {
		AuthBrowserRequest(ctx)
	}
}

func AuthApiRequest(ctx iris.Context) {
	token := ctx.GetCookie("dylank-io-auth")
	if token == "" {
		token = ctx.FormValue("token")
	}

	if token == "" {
		token = ctx.GetHeader("Authentication")
	}

	if token == "" {
		ctx.JSON(helpers.Response{
			Success:    false,
			Path:       ctx.Path(),
			Error:      helpers.ErrorMsg("Token not supplied"),
			SuccessMsg: nil,
			Data:       iris.Map{},
		})
		return
	}

	_, verify_err := services.VerifyJWT(token)

	if verify_err != nil {
		// delete cookie since its not valid
		ctx.RemoveCookie("dylank-io-auth")
		ctx.JSON(helpers.Response{
			Success:    false,
			Path:       ctx.Path(),
			Error:      helpers.ErrorMsg("Access Denied"),
			SuccessMsg: nil,
			Data:       iris.Map{},
		})
		return
	} else {
		ctx.Next()
	}
}

func AuthBrowserRequest(ctx iris.Context) {
	cookie := ctx.GetCookie("dylank-io-auth")

	if cookie == "" {
		ctx.View("auth/login.pug", iris.Map{"Redirect": ctx.Path()})
		return
	}

	_, verify_err := services.VerifyJWT(cookie)

	if verify_err != nil {
		// delete cookie since its not valid
		ctx.RemoveCookie("dylank-io-auth")
		ctx.View("auth/login.pug", iris.Map{"Redirect": ctx.Path()})
		return
	} else {
		ctx.Next()
	}
}

func isApiRequest(ctx iris.Context) bool {
	return strings.Contains(ctx.Path(), "api")
}
