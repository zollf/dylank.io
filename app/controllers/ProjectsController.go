package controllers

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/gosimple/slug"
	"github.com/kataras/iris/v12"
)

func CreateOrEditProject(ctx iris.Context) {
	if invalidProjectRequest(ctx) {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Params are invalid"), nil, iris.Map{"tag": nil})
		return
	}

	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to get or create id"), nil, iris.Map{"project": nil})
	} else {
		// Get all tags then filter
		var checked_tags []*models.Tag
		tags, _ := models.GetTags()
		input_tags := ctx.FormValues()["tags"]

		for _, tag := range tags {
			for _, input_tag := range input_tags {
				if tag.Slug == input_tag {
					checked_tags = append(checked_tags, tag)
				}
			}
		}

		project_url := ctx.FormValue("url")
		git_url := ctx.FormValue("git")

		project := &models.Project{
			ID:          id,
			Title:       ctx.FormValue("title"),
			Slug:        slug.Make(ctx.FormValue("title")),
			Description: ctx.FormValue("description"),
			Image:       "",
			URL:         &project_url,
			Git:         &git_url,
			Tags:        checked_tags,
			DateCreated: helpers.GetOrCreateDate(ctx),
			DateUpdated: time.Now().UTC().String(),
		}

		if err := models.CreateOrEditProject(project); err != nil {
			helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Failed to save project"), nil, iris.Map{"project": nil})
		} else {
			helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully saved project"), iris.Map{"project": project})
		}
	}
}

func invalidProjectRequest(ctx iris.Context) bool {
	return ctx.FormValue("title") == "" || ctx.FormValue("description") == ""
}

func DeleteProject(ctx iris.Context) {
	if ctx.FormValue("id") == "" {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg("Please include id of project"), nil, iris.Map{})
		return
	}

	if err := models.DeleteProject(ctx.FormValue("id")); err != nil {
		helpers.RedirectIfExist(ctx, helpers.ErrorMsg(err.Error()), nil, iris.Map{"project": nil})
	} else {
		helpers.RedirectIfExist(ctx, nil, helpers.SuccessMsg("Successfully deleted project"), iris.Map{})
	}
}
