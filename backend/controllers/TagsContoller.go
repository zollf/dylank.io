package controllers

import (
	"app/helpers/res"
	"app/helpers/serialize"
	"app/models/tags"
	"strconv"

	"github.com/gosimple/slug"

	"github.com/kataras/iris/v12"
)

func TagsList(ctx iris.Context) {
	if tags, cannot_list := tags.All(); cannot_list != nil {
		res.TAGS_LIST.Error(ctx, cannot_list)
	} else {
		res.TAGS_LIST.Send(ctx, iris.Map{"tags": tags})
	}
}

func TagsCreate(ctx iris.Context) {
	type Req struct {
		Title    string `json:"title" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	invalid_req, invalid_body := serialize.Body(ctx, &req)

	if invalid_req != nil {
		res.TAGS_CREATE.Error(ctx, invalid_req)
		return
	}

	if invalid_body != nil {
		res.TAGS_CREATE.ValidationError(ctx, invalid_body)
		return
	}

	tag := &tags.Tag{
		Title: req.Title,
		Slug:  slug.Make(req.Title),
	}

	if tag_create_err := tag.Create(); tag_create_err != nil {
		res.TAGS_CREATE.Error(ctx, tag_create_err)
	} else {
		res.TAGS_CREATE.Send(ctx, iris.Map{"tag": tag})
	}
}

func EditTag(ctx iris.Context) {
	type Req struct {
		ID       string `json:"id" validate:"required"`
		Title    string `json:"title" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	invalid_req, invalid_body := serialize.Body(ctx, &req)

	if invalid_req != nil {
		res.TAGS_CREATE.Error(ctx, invalid_req)
		return
	}

	if invalid_body != nil {
		res.TAGS_CREATE.ValidationError(ctx, invalid_body)
		return
	}

	id, invalid_id := strconv.ParseUint(req.ID, 10, 64)
	if invalid_id != nil {
		res.TAGS_EDIT.Error(ctx, invalid_id)
		return
	}

	tag := &tags.Tag{
		ID:    id,
		Title: req.Title,
		Slug:  slug.Make(req.Title),
	}

	if err := tag.Update(); err != nil {
		res.TAGS_EDIT.Error(ctx, err)
	} else {
		res.TAGS_EDIT.Send(ctx, iris.Map{"tag": tag})
	}
}

func DeleteTag(ctx iris.Context) {
	type Req struct {
		ID       string `json:"id" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	invalid_req, invalid_body := serialize.Body(ctx, &req)

	if invalid_req != nil {
		res.TAGS_CREATE.Error(ctx, invalid_req)
		return
	}

	if invalid_body != nil {
		res.TAGS_CREATE.ValidationError(ctx, invalid_body)
		return
	}

	tag, not_found := tags.Find(req.ID)
	if not_found != nil {
		res.TAGS_DELETE.Error(ctx, not_found)
		return
	}

	if err := tag.Delete(); err != nil {
		res.TAGS_DELETE.Error(ctx, err)
	} else {
		res.TAGS_DELETE.Send(ctx, iris.Map{})
	}
}
