package queries

import (
	"app/graphql/arguments"
	"app/graphql/interfaces"
	"app/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

func ProjectsQuery() *graphql.Field {
	return &graphql.Field{
		Args:    arguments.ProjectsArguments(),
		Type:    interfaces.ProjectsInterface(),
		Resolve: resolvers.ProjectsResolver,
	}
}
