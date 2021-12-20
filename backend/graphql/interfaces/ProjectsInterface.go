package interfaces

import (
	"github.com/graphql-go/graphql"
)

func ProjectsInterface() *graphql.List {
	return graphql.NewList(
		graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Project",
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
					"description": &graphql.Field{
						Type: graphql.String,
					},
					"image": &graphql.Field{
						Type: graphql.String,
					},
					"url": &graphql.Field{
						Type: graphql.String,
					},
					"git": &graphql.Field{
						Type: graphql.String,
					},
					"createdAt": &graphql.Field{
						Type: graphql.String,
					},
					"updatedAt": &graphql.Field{
						Type: graphql.String,
					},
					"tags": &graphql.Field{
						Type: TagsInterface(),
					},
				},
			},
		),
	)
}
