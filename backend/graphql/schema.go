package graphql

import (
	"app/utils"

	"github.com/graphql-go/graphql"
)

func GetSchema() graphql.Schema {
	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: GetQueries(),
		},
	)

	if err != nil {
		utils.Log().Error("Graphql Schema is invalid: %s", err.Error())
	}
	return schema
}
