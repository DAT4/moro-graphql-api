package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/cors"
	"log"
	"net/http"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
var h = handler.New(&handler.Config{
	Schema: &schema,
	Pretty: true,
})

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", h)
	h := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":3333", h))
}
