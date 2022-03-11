package views

import (
	"app/helpers"
	"app/models/utils"

	"github.com/kataras/iris/v12"
)

func Tags(ctx iris.Context) {
	tags, _ := utils.GetTagsData()
	helpers.RenderTemplate(ctx, "tags/tags", "admin", iris.Map{"Err": ctx.URLParam("err"), "Tags": tags})
}

func NewTag(ctx iris.Context) {
	helpers.RenderTemplate(ctx, "tags/create", "admin", iris.Map{})
}

func EditTag(ctx iris.Context) {
	id := ctx.Params().Get("id")
	tag, not_found := utils.GetTagData(id)
	if not_found != nil {
		ctx.View("404")
		return
	}

	helpers.RenderTemplate(ctx, "tags/view", "admin", iris.Map{"Tag": tag})
}
