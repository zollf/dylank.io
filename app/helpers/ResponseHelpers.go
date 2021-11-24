package helpers

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Success    bool     `json:"success"`
	Path       string   `json:"path"`
	Error      *string  `json:"error"`
	SuccessMsg *string  `json:"success_msg"`
	Data       iris.Map `json:"data"`
}

func GetOrCreateID(ctx iris.Context) (primitive.ObjectID, error) {
	if formId := ctx.FormValue("id"); formId != "" {
		return primitive.ObjectIDFromHex(formId)
	} else {
		return primitive.NewObjectID(), nil
	}
}

func GetOrCreateDate(ctx iris.Context) string {
	if dateCreated := ctx.FormValue("dateCreated"); dateCreated != "" {
		return dateCreated
	} else {
		return time.Now().UTC().String()
	}
}

func SaveRedirectIfExist(ctx iris.Context, url string, prefix string) {
	if redirect := ctx.FormValue("redirect"); redirect != "" {
		ctx.Redirect(fmt.Sprintf("%s%sredirect=%s", url, prefix, redirect))
	} else {
		ctx.Redirect(url)
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
	var success_msg *string
	success_msg = &msg
	return success_msg
}

func ErrorMsg(msg string) *string {
	var error_msg *string
	error_msg = &msg
	return error_msg
}
