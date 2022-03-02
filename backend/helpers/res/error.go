package res

import (
	"fmt"
	"net/url"

	"github.com/kataras/iris/v12"
)

func (res_type RES_TYPES) Error(ctx iris.Context, err error) {
	if redirect := ctx.FormValue("redirect"); redirect != "" {
		ctx.Redirect(fmt.Sprintf("%s?err=%s", redirect, getErrorMessage(res_type, err)))
	} else {
		ctx.StatusCode(400)
		ctx.JSON(ResJSON{
			Success:    false,
			Path:       ctx.Path(),
			Error:      String(getErrorMessage(res_type, err)),
			SuccessMsg: nil,
			Data: iris.Map{
				"error": err.Error(),
			},
		})
	}
}

func (res_type RES_TYPES) ValidationError(ctx iris.Context, validationErr []ValidationError) {
	if redirect := ctx.FormValue("redirect"); redirect != "" {
		ctx.Redirect(fmt.Sprintf("%s?err=%s", redirect, getValidationErrorMessage(res_type)))
	} else {
		ctx.StatusCode(400)
		ctx.JSON(ResJSON{
			Success:    false,
			Path:       ctx.Path(),
			Error:      String(getValidationErrorMessage(res_type)),
			SuccessMsg: nil,
			Data: iris.Map{
				"validationErr": validationErr,
			},
		})
	}
}

func getErrorMessage(res_type RES_TYPES, err error) string {
	return url.QueryEscape(fmt.Sprintf("%s", err.Error()))
}

func getValidationErrorMessage(res_type RES_TYPES) string {
	return url.QueryEscape("test")
}
