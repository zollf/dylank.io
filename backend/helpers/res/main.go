package res

import (
	"github.com/kataras/iris/v12"
)

type Response struct {
	Success bool           `json:"success"`
	Msg     string         `json:"msg"`
	Path    string         `json:"path"`
	Error   *ResponseError `json:"error"`
	Data    iris.Map       `json:"data"`
}

type ResponseError struct {
	ValidationError []ValidationError `json:"validation_error"`
	Fatal           bool              `json:"fatal"`
	Errors          []string          `json:"errors"`
}
