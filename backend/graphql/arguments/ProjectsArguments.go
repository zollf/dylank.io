package arguments

import (
	"github.com/graphql-go/graphql"
)

func ProjectsArguments() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"tags": &graphql.ArgumentConfig{
			Type:        graphql.NewList(graphql.String),
			Description: "Array of tag slug to filter by",
		},
		"offset": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "Offset the return amount, useful for pagination",
		},
		"limit": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "Limit the return amount",
		},
	}
}
