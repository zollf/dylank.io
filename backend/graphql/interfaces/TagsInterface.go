package interfaces

import (
	"github.com/graphql-go/graphql"
)

func TagsInterface() *graphql.List {
	return graphql.NewList(
		graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Tag",
				Fields: graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.String,
					},
					"slug": &graphql.Field{
						Type: graphql.String,
					},
					"title": &graphql.Field{
						Type: graphql.String,
					},
					"createdAt": &graphql.Field{
						Type: graphql.String,
					},
					"updatedAt": &graphql.Field{
						Type: graphql.String,
					},
					"count": &graphql.Field{
						Type: graphql.Int,
					},
				},
			},
		),
	)
}
