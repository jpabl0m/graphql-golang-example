package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/stobita/graphql-golang-example/internal/graphql"
	"github.com/stobita/graphql-golang-example/internal/usecase"
)

const defaultPort = "8080"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	u := usecase.New()

	resolver := graphql.NewResolver(u)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
