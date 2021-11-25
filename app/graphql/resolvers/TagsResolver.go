package resolvers

import (
	"app/models"

	"github.com/graphql-go/graphql"
)

func TagsResolver(p graphql.ResolveParams) (interface{}, error) {
	return models.GetTags()
}
