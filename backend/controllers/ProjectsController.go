package controllers

import (
	"app/helpers"
	"app/models"

	"github.com/gosimple/slug"
	"github.com/kataras/iris/v12"
)

func CreateProject(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"title", "description"}) {
		return
	}

	var checkedTags []*models.Tag
	tags, _ := models.GetTags()
	inputTags := ctx.FormValues()["tags"]

	for _, tag := range tags {
		for _, inputTag := range inputTags {
			if tag.Slug == inputTag {
				checkedTags = append(checkedTags, tag)
			}
		}
	}

	project := &models.Project{
		Title:       ctx.FormValue("title"),
		Slug:        slug.Make(ctx.FormValue("title")),
		Description: ctx.FormValue("description"),
		Image:       "",
		URL:         helpers.GetVar(ctx, "url"),
		Git:         helpers.GetVar(ctx, "git"),
		Tags:        checkedTags,
	}

	if err := models.CreateProject(project); err != nil {
		helpers.ErrorResponse(ctx, "Failed to created project", iris.Map{"error": err.Error()})
	} else {
		helpers.SuccessResponse(ctx, "Successfully created project", iris.Map{"project": project})
	}
}

func EditProject(ctx iris.Context) {
	if !helpers.ValidInputs(ctx, []string{"id", "title", "description"}) {
		return
	}

	var checkedTags []*models.Tag
	tags, _ := models.GetTags()
	inputTags := ctx.FormValues()["tags"]

	for _, tag := range tags {
		for _, inputTag := range inputTags {
			if tag.Slug == inputTag {
				checkedTags = append(checkedTags, tag)
			}
		}
	}

	project := &models.Project{
		Title:       ctx.FormValue("title"),
		Slug:        slug.Make(ctx.FormValue("title")),
		Description: ctx.FormValue("description"),
		Image:       "",
		URL:         helpers.GetVar(ctx, "url"),
		Git:         helpers.GetVar(ctx, "git"),
		Tags:        checkedTags,
	}

	if err := models.UpdateProject(project, ctx.FormValue("id")); err != nil {
		helpers.ErrorResponse(ctx, "Failed to update project", iris.Map{"error": err.Error()})
	} else {
		helpers.SuccessResponse(ctx, "Successfully updated project", iris.Map{"project": project})
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
