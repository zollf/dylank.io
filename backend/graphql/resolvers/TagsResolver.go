package resolvers

import (
	"app/models"
	"time"

	"github.com/graphql-go/graphql"
)

type TagInterface struct {
	ID        uint64    `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Count     int       `json:"count"`
}

func TagsResolver(p graphql.ResolveParams) (interface{}, error) {
	var TagsInterface []*TagInterface

	tags, tags_error := models.GetTags()

	if tags_error != nil {
		return nil, tags_error
	}

	projects, projects_error := models.GetProjects()

	if projects_error != nil {
		return nil, tags_error
	}

	for _, tag := range tags {
		TagsInterface = append(TagsInterface, &TagInterface{
			ID:        tag.ID,
			Slug:      tag.Slug,
			Title:     tag.Title,
			CreatedAt: tag.CreatedAt,
			UpdatedAt: tag.UpdatedAt,
			Count:     models.TagOccurrencesInProjects(projects, tag),
		})
	}

	return TagsInterface, tags_error
}
