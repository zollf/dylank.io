package helpers

import (
	"app/services"
	"fmt"

	"github.com/kataras/iris/v12"
)

func UploadImage(ctx iris.Context, handle string, title string) (*FileResponse, error) {
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	multip_parse_err := ctx.Request().ParseMultipartForm(maxSize)
	if multip_parse_err != nil {
		return nil, multip_parse_err
	}

	file := ctx.Request().MultipartForm.File[handle][0]

	url, s3_upload_err := services.UploadImageToS3(file, title)
	success := true

	if s3_upload_err != nil {
		return nil, s3_upload_err
	}

	if !success {
		return nil, fmt.Errorf("a file failed to upload")
	} else {
		return &FileResponse{
			Title:   title,
			Url:     url,
			Success: success,
			Error:   "",
		}, nil
	}
}
