package resolvers

import (
	"app/models"

	"github.com/graphql-go/graphql"
)

func TagsResolver(p graphql.ResolveParams) (interface{}, error) {
	var TagsInterface []*models.TagInterface

	tags, tags_error := models.GetTags()

	if tags_error != nil {
		return nil, tags_error
	}

	projects, projects_error := models.GetProjects()

	if projects_error != nil {
		return nil, tags_error
	}

	for _, tag := range tags {
		TagsInterface = append(TagsInterface, &models.TagInterface{
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
