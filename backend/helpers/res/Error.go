package res

import (
	"fmt"
	"net/url"

	"github.com/kataras/iris/v12"
)

func (res_type RES_TYPES) Error(ctx iris.Context, err error) {
	if ctx.IsStopped() {
		return
	}
	if Redirect(ctx, fmt.Sprintf("?err=%s", getErrorMessage(res_type, err))) {
		return
	}

	ctx.StatusCode(400)
	ctx.JSON(Response{
		Success: false,
		Msg:     getErrorMessage(res_type, err),
		Path:    ctx.Path(),
		Error: &ResponseError{
			Fatal: false,
		},
		Data: iris.Map{},
	})
	ctx.StopExecution()
	return
}

func (res_type RES_TYPES) ValidationError(ctx iris.Context, validationErr []ValidationError) {
	if ctx.IsStopped() {
		return
	}
	if Redirect(ctx, fmt.Sprintf("?err=%s", getValidationErrorMessage(res_type))) {
		return
	}

	ctx.StatusCode(400)
	ctx.JSON(Response{
		Success: false,
		Msg:     getValidationErrorMessage(res_type),
		Path:    ctx.Path(),
		Error: &ResponseError{
			Fatal:            false,
			ValidationErrors: validationErr,
		},
		Data: iris.Map{},
	})
	ctx.StopExecution()
	return
}

func getErrorMessage(res_type RES_TYPES, err error) string {
	return url.QueryEscape(fmt.Sprintf("%s", err.Error()))
}

func getValidationErrorMessage(res_type RES_TYPES) string {
	return url.QueryEscape("test")
}
