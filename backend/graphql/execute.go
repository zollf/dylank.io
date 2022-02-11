package graphql

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/kataras/iris/v12"
)

type RequestBody struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func ExecuteGraphqlQuery(ctx iris.Context) {
	var request RequestBody

	if err := ctx.UnmarshalBody(&request, iris.UnmarshalerFunc(json.Unmarshal)); err != nil {
		fmt.Printf("wrong result, unexpected errors: %v", err)
	}

	result := graphql.Do(graphql.Params{
		Schema:         GetSchema(),
		RequestString:  request.Query,
		VariableValues: request.Variables,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	ctx.JSON(result)
}
