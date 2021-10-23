package main

import (
	"api/graph"
	"api/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.gohtml")

	router.GET("/backend", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/backend/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	router.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	router.POST("/query", gin.WrapH(srv))

	router.Run()
}
