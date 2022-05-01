package res

import (
	"fmt"
	"net/url"

	"github.com/kataras/iris/v12"
)

func (res_type RES_TYPES) Send(ctx iris.Context, data iris.Map) {
	if Redirect(ctx, fmt.Sprintf("?success=%s", getSuccessMessage(res_type))) {
		return
	}

	ctx.StatusCode(200)
	ctx.JSON(Response{
		Success: true,
		Msg:     getSuccessMessage(res_type),
		Path:    ctx.Path(),
		Error:   nil,
		Data:    data,
	})
	return
}

func getSuccessMessage(res_type RES_TYPES) string {
	return url.QueryEscape("")
}
