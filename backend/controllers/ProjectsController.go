package controllers

import (
	"app/helpers/res"
	"app/models/assets"
	"app/models/projects"
	"app/models/tags"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/kataras/iris/v12"
)

func CreateProject(ctx iris.Context) {
	type Req struct {
		Title       string   `json:"title" validate:"required"`
		Description string   `json:"description" validate:"required"`
		Tags        []string `json:"tags"`
		Assets      []string `json:"assets"`
		Url         *string  `json:"url"`
		Git         *string  `json:"git"`
		Redirect    string   `json:"redirect"`
	}
	var req Req
	if !res.PROJECTS_CREATE.Validate(ctx, &req) {
		return
	}

	var checkedTags []*tags.Tag
	tags, _ := tags.All()
	inputTags := req.Tags

	for _, tag := range tags {
		for _, inputTag := range inputTags {
			if tag.Slug == inputTag {
				checkedTags = append(checkedTags, tag)
			}
		}
	}

	var checkAssets []*assets.Asset
	assets, _ := assets.All()
	inputAssets := req.Assets

	for _, asset := range assets {
		for _, inputAsset := range inputAssets {
			if asset.Slug == inputAsset {
				checkAssets = append(checkAssets, asset)
			}
		}
	}

	project := &projects.Project{
		Title:       req.Title,
		Slug:        slug.Make(req.Title),
		Description: req.Description,
		Assets:      checkAssets,
		URL:         req.Url,
		Git:         req.Git,
		Tags:        checkedTags,
	}

	if err := project.Create(); err != nil {
		res.PROJECTS_CREATE.Error(ctx, err)
	} else {
		res.PROJECTS_CREATE.Send(ctx, iris.Map{"project": project})
	}
}

func EditProject(ctx iris.Context) {
	type Req struct {
		ID          string   `json:"id" validate:"required"`
		Title       string   `json:"title" validate:"required"`
		Description string   `json:"description" validate:"required"`
		Tags        []string `json:"tags"`
		Assets      []string `json:"assets"`
		Url         *string  `json:"url"`
		Git         *string  `json:"git"`
		Redirect    string   `json:"redirect"`
	}
	var req Req
	if !res.PROJECTS_EDIT.Validate(ctx, &req) {
		return
	}

	id, invalid_id := strconv.ParseUint(req.ID, 10, 64)
	if invalid_id != nil {
		res.PROJECTS_EDIT.Error(ctx, invalid_id)
		return
	}

	var checkedTags []*tags.Tag
	tags, _ := tags.All()
	inputTags := req.Tags

	for _, tag := range tags {
		for _, inputTag := range inputTags {
			if tag.Slug == inputTag {
				checkedTags = append(checkedTags, tag)
			}
		}
	}

	var checkAssets []*assets.Asset
	assets, _ := assets.All()
	inputAssets := req.Assets

	for _, asset := range assets {
		for _, inputAsset := range inputAssets {
			if asset.Slug == inputAsset {
				checkAssets = append(checkAssets, asset)
			}
		}
	}

	project := &projects.Project{
		ID:          id,
		Title:       req.Title,
		Slug:        slug.Make(req.Title),
		Description: req.Description,
		Assets:      checkAssets,
		URL:         req.Url,
		Git:         req.Git,
		Tags:        checkedTags,
	}

	if err := project.Update(); err != nil {
		res.PROJECTS_EDIT.Error(ctx, err)
	} else {
		res.PROJECTS_EDIT.Send(ctx, iris.Map{"project": project})
	}
}

func DeleteProject(ctx iris.Context) {
	type Req struct {
		ID       string `json:"id" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	if !res.PROJECTS_DELETE.Validate(ctx, &req) {
		return
	}

	project, not_found := projects.Find(req.ID)
	if not_found != nil {
		res.PROJECTS_DELETE.Error(ctx, not_found)
		return
	}

	if err := project.Delete(); err != nil {
		res.PROJECTS_DELETE.Error(ctx, err)
	} else {
		res.PROJECTS_DELETE.Send(ctx, iris.Map{})
	}
}

func ListProject(ctx iris.Context) {
	projects, err := projects.All()
	if err != nil {
		res.PROJECTS_LIST.Error(ctx, err)
		return
	}
	res.PROJECTS_LIST.Send(ctx, iris.Map{"projects": projects})
}
