package middleware

import (
	"app/models"
	"app/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

func AuthOpts() basicauth.Options {
	// I use id rather than username
	allowFunc := func(ctx iris.Context, username string, password string) (interface{}, bool) {
		user, err := models.GetUserWithPassword(username, password)
		return user, err == nil
	}

	opts := basicauth.Options{
		Realm:        basicauth.DefaultRealm,
		ErrorHandler: basicauth.DefaultErrorHandler,
		Allow:        allowFunc,
	}

	return opts
}

func AuthRequired(ctx iris.Context) {
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
