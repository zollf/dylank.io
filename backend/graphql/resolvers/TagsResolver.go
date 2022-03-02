package resolvers

import (
	"app/models/projects"
	"app/models/tags"
	"app/models/utils"

	"github.com/graphql-go/graphql"
)

func TagsResolver(p graphql.ResolveParams) (interface{}, error) {
	var TagsInterface []*utils.TagInterface

	tags, tags_error := tags.All()

	if tags_error != nil {
		return nil, tags_error
	}

	projects, projects_error := projects.All()

	if projects_error != nil {
		return nil, tags_error
	}

	for _, tag := range tags {
		TagsInterface = append(TagsInterface, &utils.TagInterface{
			ID:        tag.ID,
			Slug:      tag.Slug,
			Title:     tag.Title,
			CreatedAt: tag.CreatedAt,
			UpdatedAt: tag.UpdatedAt,
			Count:     utils.TagOccurrencesInProjects(projects, tag),
		})
	}

	return TagsInterface, tags_error
}
