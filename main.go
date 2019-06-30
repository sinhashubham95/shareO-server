package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/sinhashubham95/shareO-server/graphql"
)

func main() {
	r := gin.Default()
	r.Any("/graphql", graphqlHandler())
	r.Run(":8080")
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:     graphql.GetSchema(),
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
