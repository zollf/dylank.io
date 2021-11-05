package controllers

import (
	"app/models"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrEditProject(ctx iris.Context) {
	var ID primitive.ObjectID
	if ctx.FormValue("id") != "" {
		pid, pid_err := primitive.ObjectIDFromHex(ctx.FormValue("id"))

		if pid_err != nil {
			ctx.Redirect(fmt.Sprintf("/admin/projects?err=%s", pid_err.Error()))
			return
		}

		ID = pid
	} else {
		ID = primitive.NewObjectID()
	}

	var DateCreated string
	if ctx.FormValue("dateCreated") != "" {
		DateCreated = ctx.FormValue("dateCreated")
	} else {
		DateCreated = time.Now().UTC().String()
	}

	var project_url string
	if ctx.FormValue("url") != "" {
		project_url = ctx.FormValue("url")
	}

	var git_url string
	if ctx.FormValue("git") != "" {
		git_url = ctx.FormValue("git")
	}

	project := &models.Project{
		ID:          ID,
		Title:       ctx.FormValue("title"),
		Slug:        ctx.FormValue("slug"),
		Description: ctx.FormValue("description"),
		Image:       "",
		URL:         &project_url,
		Git:         &git_url,
		DateCreated: DateCreated,
		DateUpdated: time.Now().UTC().String(),
	}

	err := models.CreateOrEditProject(project)

	if err != nil {
		ctx.Redirect(fmt.Sprintf("/admin/projects?err=%s", err.Error()))
		return
	} else {
		ctx.Redirect("/admin/projects?success=true")
		return
	}
}
