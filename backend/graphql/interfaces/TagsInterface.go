package interfaces

import (
	"app/graphql/utils"

	"github.com/graphql-go/graphql"
)

var _TagsInterface *graphql.List

func TagsInterface() *graphql.List {
	return utils.CreateList("Tag", graphql.NewList(
		graphql.NewObject(
			graphql.ObjectConfig{
				Name:        "Tag",
				Description: "Singular Tag that can describe a langauge a project is written in.",
				Fields: graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.NewNonNull(graphql.String),
					},
					"slug": &graphql.Field{
						Type: graphql.NewNonNull(graphql.String),
					},
					"title": &graphql.Field{
						Type: graphql.NewNonNull(graphql.String),
					},
					"createdAt": &graphql.Field{
						Type: graphql.NewNonNull(graphql.String),
					},
					"updatedAt": &graphql.Field{
						Type: graphql.NewNonNull(graphql.String),
					},
					"count": &graphql.Field{
						Type: graphql.Int,
					},
				},
			},
		),
	),
	)
}
