package config

import (
	"app/graph"
	"app/graph/generated"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func Graphql(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	path := r.URL.Path
	if strings.HasPrefix(path, "/admin/playground") {
		playground.Handler("GraphQL playground", "/api/graphql").ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(path, "/api/graphql") {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
		srv.ServeHTTP(w, r)
		return
	}

	router.ServeHTTP(w, r)
}
