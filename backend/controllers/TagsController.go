package controllers

import (
	"app/helpers/res"
	"app/models/tags"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/kataras/iris/v12"
)

type TagsController struct{}

// Get Tag
// Method:   GET
// Resource: /api/tags
func (c *TagsController) Get(ctx iris.Context) {
	if tags, cannot_list := tags.All(); cannot_list != nil {
		res.TAGS_LIST.Error(ctx, cannot_list)
	} else {
		res.TAGS_LIST.Send(ctx, iris.Map{"tags": tags})
	}
}

// Get Tag by ID
// Method:   GET
// Resource: /api/tags/{id:int}
func (c *TagsController) GetBy(id int, ctx iris.Context) {
	tag, cannot_find := tags.Find(id)
	if cannot_find != nil {
		res.TAG_GET.Error(ctx, cannot_find)
	}
	res.TAG_GET.Send(ctx, iris.Map{"tag": tag})
}

// Create Tag
// Method:   POST
// Resource: /api/tags/create
func (c *TagsController) PostCreate(req tags.TagCreateRequest, ctx iris.Context) {
	res.TAG_CREATE.Validate(ctx, &req)
	if ctx.IsStopped() {
		return
	}

	tag := &tags.Tag{
		Title: req.Title,
		Slug:  slug.Make(req.Title),
	}

	if tag_create_err := tag.Create(); tag_create_err != nil {
		res.TAG_CREATE.Error(ctx, tag_create_err)
	}

	res.TAG_CREATE.Send(ctx, iris.Map{"tag": tag})
}

// Edit Tag
// Method:   POST
// Resource: http://localhost/api/tag
func (c *TagsController) PostEdit(req tags.TagEditRequest, ctx iris.Context) {
	res.TAG_EDIT.Validate(ctx, &req)
	if ctx.IsStopped() {
		return
	}

	id, invalid_id := strconv.ParseUint(req.ID, 10, 64)

	if invalid_id != nil {
		res.TAG_EDIT.Error(ctx, invalid_id)
	}

	tag := &tags.Tag{
		ID:    id,
		Title: req.Title,
		Slug:  slug.Make(req.Title),
	}

	if err := tag.Update(); err != nil {
		res.TAG_EDIT.Error(ctx, err)
	}

	res.TAG_EDIT.Send(ctx, iris.Map{"tag": tag})
}

// Delete Tag
// Method:   POST
// Resource: http://localhost/api/tag
func (c *TagsController) PostDelete(req tags.TagDeleteRequest, ctx iris.Context) {
	res.TAG_DELETE.Validate(ctx, &req)
	if ctx.IsStopped() {
		return
	}

	tag, not_found := tags.Find(req.ID)
	if not_found != nil {
		res.TAG_DELETE.Error(ctx, not_found)
	}

	if err := tag.Delete(); err != nil {
		res.TAG_DELETE.Error(ctx, err)
	}

	res.TAG_DELETE.Send(ctx, iris.Map{})
}
