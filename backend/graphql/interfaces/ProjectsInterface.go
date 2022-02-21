package interfaces

import (
	"app/graphql/utils"

	"github.com/graphql-go/graphql"
)

func ProjectsInterface() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "Projects",
			Description: "All projects",
			Fields: graphql.Fields{
				"items": &graphql.Field{
					Type: ProjectItemsInterface(),
				},
				"tags": &graphql.Field{
					Type: TagsInterface(),
				},
				"items_total": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"total": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
		},
	)
}

func ProjectItemsInterface() *graphql.List {
	return utils.CreateList("ProjectItems", graphql.NewList(ProjectInterface()))
}

func ProjectInterface() *graphql.Object {
	return utils.CreateObject("Project", graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Project",
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
				"description": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"assets": &graphql.Field{
					Type: AssetsInterface(),
				},
				"url": &graphql.Field{
					Type: graphql.String,
				},
				"git": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"updatedAt": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"tags": &graphql.Field{
					Type: TagsInterface(),
				},
			},
		},
	),
	)
}
