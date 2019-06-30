package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
	r := gin.Default()
	r.Any("/graphql", h)
	r.Run(":8080")
}
