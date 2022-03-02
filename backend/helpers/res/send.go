package res

import (
	"fmt"
	"net/url"

	"github.com/kataras/iris/v12"
)

func (res_type RES_TYPES) Send(ctx iris.Context, data iris.Map) {
	if redirect := ctx.FormValue("redirect"); redirect != "" {
		ctx.Redirect(fmt.Sprintf("%s?success=%s", redirect, getSuccessMessage(res_type)))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(ResJSON{
			Success:    true,
			Path:       ctx.Path(),
			Error:      String(getSuccessMessage(res_type)),
			SuccessMsg: nil,
			Data:       data,
		})
	}
}

func getSuccessMessage(res_type RES_TYPES) string {
	return url.QueryEscape("")
}
