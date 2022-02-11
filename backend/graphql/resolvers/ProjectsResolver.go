package resolvers

import (
	"app/database"
	"app/models"

	"github.com/graphql-go/graphql"
)

type ProjectsResolverType struct {
	Items      []*models.Project      `json:"items"`
	Tags       []*models.TagInterface `json:"tags"`
	Total      int64                  `json:"total"`
	ItemsTotal int                    `json:"items_total"`
}

func ProjectsResolver(p graphql.ResolveParams) (interface{}, error) {
	var projects []*models.Project

	if db, err := database.Open(); err == nil {
		var total int64
		db.Model(&models.Project{}).Count(&total)
		query := db.Preload("Tags")

		if tags, ok := p.Args["tags"].([]interface{}); ok {
			if len(tags) > 0 && tags[0] != "all" {
				query.
					Joins("JOIN project_tags on project_tags.project_id = projects.id").
					Joins("JOIN tags on project_tags.tag_id = tags.id").
					Where("tags.slug IN ?", tags).
					Group("projects.id").
					Having("COUNT(DISTINCT tags.slug) = ?", len(tags))
			}
		}

		query.Find(&projects)
		items_total := len(projects)
		if query.Error != nil {
			return nil, query.Error
		}

		tags := models.TagsOccurrencesInProjects(projects)

		if limit, ok := p.Args["limit"].(int); ok {
			query.Limit(limit)
		}

		if offset, ok := p.Args["offset"].(int); ok {
			query.Offset(offset)
		}

		query.Find(&projects)

		result := ProjectsResolverType{
			Items:      projects,
			Tags:       tags,
			Total:      total,
			ItemsTotal: items_total,
		}
		return result, nil
	} else {
		return nil, err
	}
}
