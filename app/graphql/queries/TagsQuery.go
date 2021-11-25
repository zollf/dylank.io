package queries

import (
	"app/graphql/arguments"
	"app/graphql/interfaces"
	"app/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

func TagsQuery() *graphql.Field {
	return &graphql.Field{
		Args:    arguments.TagsArguments(),
		Type:    interfaces.TagsInterface(),
		Resolve: resolvers.TagsResolver,
	}
}
