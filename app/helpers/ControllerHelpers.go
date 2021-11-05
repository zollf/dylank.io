package helpers

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func RedirectIfExist(ctx iris.Context, defaultUrl string) {
	if redirect := ctx.FormValue("redirect"); redirect != "" {
		ctx.Redirect(redirect)
	} else {
		ctx.Redirect(defaultUrl)
	}
}
