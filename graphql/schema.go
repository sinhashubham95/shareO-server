package graphql

import (
	"github.com/graphql-go/graphql"
	"log"
)

var schema graphql.Schema

func init() {
	var err error
	schema, err = graphql.NewSchema(graphql.SchemaConfig{})
	if err != nil {
		log.Fatal("Error initializing schema.")
	}
}
