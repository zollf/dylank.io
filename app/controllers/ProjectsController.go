package controllers

import (
	"app/helpers"
	"app/models"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

func CreateOrEditProject(ctx iris.Context) {
	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/projects?err=%s", id_err.Error()))
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
			Slug:        ctx.FormValue("slug"),
			Description: ctx.FormValue("description"),
			Image:       "",
			URL:         &project_url,
			Git:         &git_url,
			Tags:        checked_tags,
			DateCreated: helpers.GetOrCreateDate(ctx),
			DateUpdated: time.Now().UTC().String(),
		}

		if err := models.CreateOrEditProject(project); err != nil {
			ctx.Redirect(fmt.Sprintf("/admin/projects?err=%s", err.Error()))
		} else {
			ctx.Redirect("/admin/projects?success=true")
		}
	}
}

func DeleteProject(ctx iris.Context) {
	if err := models.DeleteProject(ctx.FormValue("id")); err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/projects?err=%s", err.Error()))
	} else {
		ctx.Redirect("/admin/projects?success=true")
	}
}
