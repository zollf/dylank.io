package views

import (
	"app/helpers"
	"app/models"

	"github.com/kataras/iris/v12"
)

func Projects(ctx iris.Context) {
	projects, _ := models.GetProjectData()
	helpers.RenderTemplate(ctx, "projects/projects", "admin", iris.Map{"Projects": projects, "Err": ctx.URLParam("err")})
}

func NewProject(ctx iris.Context) {
	tags, _ := models.GetTagsData()
	assets, _ := models.GetAssetsData()
	helpers.RenderTemplate(ctx, "projects/create", "admin", iris.Map{"Tags": tags, "Assets": assets})
}

func EditProject(ctx iris.Context) {
	id := ctx.Params().Get("id")
	project, tags, assets, not_found := models.GetProjectDataWithTagAndAsset(id)
	if not_found != nil {
		ctx.View("404")
		return
	}
	helpers.RenderTemplate(ctx, "projects/view", "admin", iris.Map{"Project": project, "Tags": tags, "Assets": assets})
}
