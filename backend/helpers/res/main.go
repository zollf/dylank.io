package res

import (
	"github.com/kataras/iris/v12"
)

type Res struct {
	Err  error
	Data interface{}
}

type ResJSON struct {
	Success    bool     `json:"success"`
	Path       string   `json:"path"`
	Error      *string  `json:"error"`
	SuccessMsg *string  `json:"success_msg"`
	Data       iris.Map `json:"data"`
}
