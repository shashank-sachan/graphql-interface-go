package main

import (
	"graphql-interface-go/generated"
	"graphql-interface-go/resolvers"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const serverAddress = "0.0.0.0:9090"

func graphqlHandler() gin.HandlerFunc {
	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.Use(extension.Introspection{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	r.GET("/", playgroundHandler())
	r.GET("/graphql", graphqlHandler())
	r.POST("/graphql", graphqlHandler())

	log.Printf("connect to %s for GraphQL playground", serverAddress)
	r.Run(serverAddress)
}
