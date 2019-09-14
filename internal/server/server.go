package server

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/stobita/graphql-golang-example/internal/graphql"
	"github.com/stobita/graphql-golang-example/internal/storage"
	"github.com/stobita/graphql-golang-example/internal/usecase"
)

const defaultPort = "8080"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()

	redisClient := storage.NewRedisClient()
	sessionStore := storage.NewSessionStore(redisClient)
	middleware := newMiddleware(sessionStore)
	r.Use(middleware.getAuthenticationMiddleware())

	// r.Use(authenticationMiddleware())
	r.GET("/", rootHandler())
	r.POST("/query", queryHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run(":" + port))
}

func queryHandler() gin.HandlerFunc {
	u := usecase.New()
	resolver := graphql.NewResolver(u)
	c := graphql.Config{Resolvers: resolver}
	c.Directives.RequireAuth = graphql.RequireAuthDirective
	h := handler.GraphQL(graphql.NewExecutableSchema(c))
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
