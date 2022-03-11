package res

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type ValidationError struct {
	ActualTag string `json:"tag"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	Param     string `json:"param"`
}

var validate *validator.Validate

func WrapValidationErrors(errs validator.ValidationErrors) []ValidationError {
	validationErrors := make([]ValidationError, 0, len(errs))
	for _, validationErr := range errs {
		validationErrors = append(validationErrors, ValidationError{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Kind:      validationErr.Kind().String(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})
	}

	return validationErrors
}

func (res_type RES_TYPES) Validate(ctx iris.Context, outPtr interface{}) bool {
	body_err := ctx.ReadBody(outPtr)
	if body_err != nil {
		res_type.Error(ctx, body_err)
		return false
	}

	validate = validator.New()

	invalid_req := validate.Struct(outPtr)
	if invalid_req != nil {
		res_type.ValidationError(ctx, WrapValidationErrors(invalid_req.(validator.ValidationErrors)))
		return false
	}

	return true
}
