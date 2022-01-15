package helpers

import (
	"app/services"
	"fmt"

	"github.com/kataras/iris/v12"
)

func UploadImage(ctx iris.Context, name string) ([]*FileResponse, error) {
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	mp_err := ctx.Request().ParseMultipartForm(maxSize)
	if mp_err != nil {
		return nil, mp_err
	}

	form := ctx.Request().MultipartForm
	files := form.File[name]
	var uploadedFiles []*FileResponse

	file := files[0]

	title := ctx.FormValue("title")

	url, err := services.UploadImageToS3(file, title)
	success := true
	errorMsg := ""

	if err != nil {
		success = false
		errorMsg = err.Error()
	}

	uploadedFiles = append(uploadedFiles, &FileResponse{
		Title:   title,
		Url:     url,
		Success: success,
		Error:   errorMsg,
	})

	if !success {
		return uploadedFiles, fmt.Errorf("a file failed to upload")
	} else {
		return uploadedFiles, nil
	}
}
