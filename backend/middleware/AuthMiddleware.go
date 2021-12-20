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
		authApiRequest(ctx)
	} else {
		authBrowserRequest(ctx)
	}
}

func authApiRequest(ctx iris.Context) {
	token := helpers.GetToken(ctx)

	if token == "" {
		ctx.StatusCode(400)
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
		ctx.StatusCode(401)

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

func authBrowserRequest(ctx iris.Context) {
	cookie := ctx.GetCookie("dylank-io-auth")
	err := ctx.URLParam("err")
	success := ctx.URLParam("success")

	if cookie == "" {
		helpers.RenderTemplate(ctx, "auth/login", "base", iris.Map{
			"Redirect": ctx.Path(),
			"Err":      err,
			"Success":  success,
		})
		return
	}

	_, verify_err := services.VerifyJWT(cookie)

	if verify_err != nil {
		// delete cookie since its not valid
		ctx.RemoveCookie("dylank-io-auth")
		helpers.RenderTemplate(ctx, "auth/login", "base", iris.Map{
			"Redirect": ctx.Path(),
			"Err":      err,
			"Success":  success,
		})
		return
	} else {
		ctx.Next()
	}
}

func isApiRequest(ctx iris.Context) bool {
	return strings.Contains(ctx.Path(), "api")
}
