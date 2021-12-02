package graphql

import (
	"app/graphql/queries"

	"github.com/graphql-go/graphql"
)

func GetQueries() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"ping":     queries.PingPongQuery(),
				"projects": queries.ProjectsQuery(),
				"tags":     queries.TagsQuery(),
			},
		},
	)
}
