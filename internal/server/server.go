package server

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/stobita/graphql-golang-example/internal/graphql"
	"github.com/stobita/graphql-golang-example/internal/usecase"
)

const defaultPort = "8080"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()

	r.POST("/", rootHandler())
	r.GET("/query", queryHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run(":" + port))
}

func queryHandler() gin.HandlerFunc {
	u := usecase.New()
	resolver := graphql.NewResolver(u)
	h := handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func rootHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
