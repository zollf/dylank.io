package interfaces

import (
	"github.com/graphql-go/graphql"
)

func PingPongInterface() *graphql.Scalar {
	return graphql.String
}
