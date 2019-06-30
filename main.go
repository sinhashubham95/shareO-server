package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/sinhashubham95/shareO-server/graphql"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:     graphql.GetSchema(),
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
	r := gin.Default()
	r.Any("/graphql", h)
	r.Run(":8080")
}
