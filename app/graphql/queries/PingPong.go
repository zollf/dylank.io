package queries

import (
	"app/graphql/arguments"
	"app/graphql/interfaces"
	"app/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

func PingPongQuery() *graphql.Field {
	return &graphql.Field{
		Args:    arguments.PingPongArguments(),
		Type:    interfaces.PingPongInterface(),
		Resolve: resolvers.PingPongResolver,
	}
}
