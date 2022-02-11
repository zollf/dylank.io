package graphql

import (
	"log"

	"github.com/graphql-go/graphql"
)

func GetSchema() graphql.Schema {
	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: GetQueries(),
		},
	)

	if err != nil {
		log.Printf("GRAPHQL SCHEMA ERROR: %s", err.Error())
	}
	return schema
}
