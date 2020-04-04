package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"queries"
	"mutations"
	"./handlers"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queries.RootQuery,
	Mutation: mutations.RootMutation,
})

func main() {
	r := handlers.Router()

	fmt.Println("Starting server on the 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
