package serialize

import (
	"app/helpers/res"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

var validate *validator.Validate

func Body(ctx iris.Context, outPtr interface{}) (error, []res.ValidationError) {
	body_err := ctx.ReadBody(outPtr)
	if body_err != nil {
		return body_err, nil
	}

	validate = validator.New()

	invalid_req := validate.Struct(outPtr)
	if invalid_req != nil {
		return nil, res.WrapValidationErrors(invalid_req.(validator.ValidationErrors))
	}
	return nil, nil
}
