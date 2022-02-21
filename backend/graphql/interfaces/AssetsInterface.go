package interfaces

import (
	"app/graphql/utils"

	"github.com/graphql-go/graphql"
)

func AssetsInterface() *graphql.List {
	return utils.CreateList("Asset", graphql.NewList(
		graphql.NewObject(
			graphql.ObjectConfig{
				Name:        "Asset",
				Description: "Asset",
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
					"url": &graphql.Field{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			},
		),
	),
	)
}
