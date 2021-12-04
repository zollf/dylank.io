package views

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/kataras/iris/v12"
)

func Tags(ctx iris.Context) {
	err := ctx.URLParam("err")
	type TagData struct {
		ID          string
		Index       int
		Title       string
		Slug        string
		DateCreated string
		DateUpdated string
	}

	tags, t_err := models.GetTags()
	if t_err != nil {
		err = t_err.Error()
	}

	var tag_data []*TagData
	layout := "2006-01-02 15:04:05 -0700 MST"
	zone, _ := time.LoadLocation("Australia/Perth")

	for i, tag := range tags {
		DateCreated, _ := time.Parse(layout, tag.DateCreated)
		DateUpdated, _ := time.Parse(layout, tag.DateUpdated)

		tag_data = append(tag_data, &TagData{
			ID:          tag.ID.Hex(),
			Index:       i + 1,
			Title:       tag.Title,
			Slug:        tag.Slug,
			DateCreated: DateCreated.In(zone).Format(time.RFC822),
			DateUpdated: DateUpdated.In(zone).Format(time.RFC822),
		})
	}

	helpers.RenderTemplate(ctx, "tags/tags", "admin", iris.Map{"Err": err, "Tags": tag_data})
}

func NewTag(ctx iris.Context) {
	helpers.RenderTemplate(ctx, "tags/create", "admin", iris.Map{})
}

func EditTag(ctx iris.Context) {
	type TagData struct {
		ID          string
		Index       int
		Title       string
		Slug        string
		DateCreated string
		DateUpdated string
	}

	id := ctx.Params().Get("id")
	tag, not_found := models.GetTag(id)

	if not_found != nil {
		ctx.View("404")
		return
	}

	tag_data := TagData{
		ID:          tag.ID.Hex(),
		Title:       tag.Title,
		Slug:        tag.Slug,
		DateCreated: tag.DateCreated,
	}

	helpers.RenderTemplate(ctx, "tags/view", "admin", iris.Map{"Tag": tag_data})
}
