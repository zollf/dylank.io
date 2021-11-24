package controllers

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/gosimple/slug"
	"github.com/kataras/iris/v12"
)

func CreateOrEditProject(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"title", "description"}) {
		return
	}

	if id, id_err := helpers.GetOrCreateID(ctx); id_err != nil {
		helpers.ErrorResponse(ctx, "Failed to get or create id", iris.Map{"project": nil, "error": id_err.Error()})
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

		project := &models.Project{
			ID:          id,
			Title:       ctx.FormValue("title"),
			Slug:        slug.Make(ctx.FormValue("title")),
			Description: ctx.FormValue("description"),
			Image:       "",
			URL:         helpers.GetVar(ctx, "url"),
			Git:         helpers.GetVar(ctx, "git"),
			Tags:        checked_tags,
			DateCreated: helpers.GetOrCreateDate(ctx),
			DateUpdated: time.Now().UTC().String(),
		}

		if err := models.CreateOrEditProject(project); err != nil {
			helpers.ErrorResponse(ctx, "Failed to save project", iris.Map{"project": nil, "error": err.Error()})
		} else {
			helpers.SuccessResponse(ctx, "Successfully saved project", iris.Map{"project": project})
		}
	}
}

func DeleteProject(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"id"}) {
		return
	}

	if err := models.DeleteProject(ctx.FormValue("id")); err != nil {
		helpers.ErrorResponse(ctx, "Failed to delete project", iris.Map{"error": err.Error()})
	} else {
		helpers.SuccessResponse(ctx, "Successfully deleted project", iris.Map{})
	}
}

func ListProject(ctx iris.Context) {
	projects, err := models.GetProjects()
	if err != nil {
		helpers.ErrorResponse(ctx, "Failed to list projects", iris.Map{"error": err.Error()})
	}
	helpers.SuccessResponse(ctx, "Successfully listed projects", iris.Map{"projects": projects})
}
