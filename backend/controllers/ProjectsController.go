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

type ProjectsController struct{}

// Get Projects
// Method:   GET
// Resource: /api/projects
func (c *ProjectsController) Get(ctx iris.Context) {
	projects, err := projects.All()
	if err != nil {
		res.PROJECTS_LIST.Error(ctx, err)
	}
	res.PROJECTS_LIST.Send(ctx, iris.Map{"projects": projects})
}

// Get Project by ID
// Method:   GET
// Resource: /api/projects/{id:int}
func (c *ProjectsController) GetBy(id int, ctx iris.Context) {
	project, cannot_find := projects.Find(id)
	if cannot_find != nil {
		res.PROJECT_GET.Error(ctx, cannot_find)
	}
	res.PROJECT_GET.Send(ctx, iris.Map{"project": project})
}

// Create Project
// Method:   POST
// Resource: /api/tags/create
func (c *ProjectsController) PostCreate(req projects.ProjectCreateRequest, ctx iris.Context) {
	res.PROJECT_CREATE.Validate(ctx, &req)
	if ctx.IsStopped() {
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
		res.PROJECT_CREATE.Error(ctx, err)
	} else {
		res.PROJECT_CREATE.Send(ctx, iris.Map{"project": project})
	}
}

// Edit Project
// Method:   POST
// Resource: /api/projects/create
func (c *ProjectsController) PostEdit(req projects.ProjectEditRequest, ctx iris.Context) {
	res.PROJECT_EDIT.Validate(ctx, &req)
	if ctx.IsStopped() {
		return
	}

	id, invalid_id := strconv.ParseUint(req.ID, 10, 64)
	if invalid_id != nil {
		res.PROJECT_EDIT.Error(ctx, invalid_id)
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
		res.PROJECT_EDIT.Error(ctx, err)
	} else {
		res.PROJECT_EDIT.Send(ctx, iris.Map{"project": project})
	}
}

// Delete Project
// Method:   POST
// Resource: /api/projects/delete
func (c *ProjectsController) PostDelete(req projects.ProjectDeleteRequest, ctx iris.Context) {
	res.PROJECT_DELETE.Validate(ctx, &req)
	if ctx.IsStopped() {
		return
	}

	project, not_found := projects.Find(req.ID)
	if not_found != nil {
		res.PROJECT_DELETE.Error(ctx, not_found)
		return
	}

	if err := project.Delete(); err != nil {
		res.PROJECT_DELETE.Error(ctx, err)
	} else {
		res.PROJECT_DELETE.Send(ctx, iris.Map{})
	}
}
