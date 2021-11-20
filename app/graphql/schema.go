package graphql

import (
	"github.com/graphql-go/graphql"
)

func GetSchema() graphql.Schema {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: GetQueries(),
		},
	)
	return schema
}
