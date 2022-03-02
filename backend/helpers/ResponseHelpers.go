package helpers

import (
	"app/services"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

type Response struct {
	Success    bool     `json:"success"`
	Path       string   `json:"path"`
	Error      *string  `json:"error"`
	SuccessMsg *string  `json:"success_msg"`
	Data       iris.Map `json:"data"`
}

func GetOrCreateDate(ctx iris.Context) string {
	if dateCreated := ctx.FormValue("dateCreated"); dateCreated != "" {
		return dateCreated
	} else {
		return time.Now().UTC().String()
	}
}

func RedirectIfExist(ctx iris.Context, errMsg *string, successMsg *string, data iris.Map) {
	if redirect := ctx.FormValue("redirect"); redirect != "" {
		if errMsg != nil {
			ctx.Redirect(fmt.Sprintf("%s?err=%s", redirect, url.QueryEscape(*errMsg)))
			return
		}
		if successMsg != nil {
			ctx.Redirect(fmt.Sprintf("%s?success=%s", redirect, url.QueryEscape(*successMsg)))
			return
		}
		ctx.Redirect(fmt.Sprintf("%s?success=%s", redirect, "Action was successful"))
	} else {
		success := true
		if errMsg != nil {
			success = false
			ctx.StatusCode(400)
		}
		ctx.JSON(Response{
			Success:    success,
			Path:       ctx.Path(),
			Error:      errMsg,
			SuccessMsg: successMsg,
			Data:       data,
		})
	}
}

func SuccessMsg(msg string) *string {
	success_msg := &msg
	return success_msg
}

func ErrorMsg(msg string) *string {
	error_msg := &msg
	return error_msg
}

func ErrorResponse(ctx iris.Context, err string, data iris.Map) {
	RedirectIfExist(ctx, ErrorMsg(err), nil, data)
}

func SuccessResponse(ctx iris.Context, msg string, data iris.Map) {
	RedirectIfExist(ctx, nil, SuccessMsg(msg), data)
}

func GetVar(ctx iris.Context, val string) *string {
	var result *string
	if formValue := ctx.FormValue(val); formValue != "" {
		result = &formValue
	}

	return result
}

func ValidInputs(ctx iris.Context, inputs []string) bool {
	errorMap := make(map[string]string)
	valid := true
	missingVars := ""
	for _, input := range inputs {
		if ctx.FormValue(input) == "" {
			errorMap[input] = fmt.Sprintf("Paramater %s is missing or empty", input)
			missingVars += fmt.Sprintf("%s ", input)
			valid = false
		}
	}

	if !valid {
		ErrorResponse(ctx, "Parameter(s) are missing or empty", iris.Map{"errors": errorMap})
	}

	return valid
}

type FileResponse struct {
	Title   string
	Url     string
	Success bool
	Error   string
}

func UploadImages(ctx iris.Context, name string) ([]*FileResponse, error) {
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	mp_err := ctx.Request().ParseMultipartForm(maxSize)
	if mp_err != nil {
		return nil, mp_err
	}

	form := ctx.Request().MultipartForm
	files := form.File[name]
	failures := 0
	var uploadedFiles []*FileResponse

	for _, file := range files {
		title := strings.Split(file.Filename, ".")[0]

		url, err := services.UploadImageToS3(file, file.Filename)
		success := true
		errorMsg := ""

		if err != nil {
			success = false
			failures = failures + 1
			errorMsg = err.Error()
		}

		uploadedFiles = append(uploadedFiles, &FileResponse{
			Title:   title,
			Url:     url,
			Success: success,
			Error:   errorMsg,
		})
	}

	if failures != 0 {
		return uploadedFiles, fmt.Errorf("a file failed to upload")
	} else {
		return uploadedFiles, nil
	}
}
