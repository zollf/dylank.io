package resolvers

import (
	"app/models"

	"github.com/graphql-go/graphql"
)

func ProjectsResolver(p graphql.ResolveParams) (interface{}, error) {
	return models.GetProjects()
}
