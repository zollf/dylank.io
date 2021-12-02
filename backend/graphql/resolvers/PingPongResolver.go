package resolvers

import "github.com/graphql-go/graphql"

func PingPongResolver(p graphql.ResolveParams) (interface{}, error) {
	return "pong", nil
}
