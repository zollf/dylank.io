package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func GetSchema() graphql.Schema {
	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: GetQueries(),
		},
	)

	if err != nil {
		fmt.Printf("Graphql Schema is invalid: %s", err.Error())
	}
	return schema
}
