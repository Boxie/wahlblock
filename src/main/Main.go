package main

import (

	"net/http"
	"github.com/graphql-go/graphql"
	graph "github.com/boxie/wahlblock/src/main/graphql"

	"github.com/boxie/wahlblock/src/main/config"
	"github.com/graphql-go/handler"
)



func main() {

	config.MigrateDatabase()

	//Assign graphql to http handler
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: graph.RootQuery,
		Mutation: graph.RootMutation,
	})


	graphql := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})


	fs := http.FileServer(http.Dir("static"))

	http.Handle("/graphql", graphql)
	// static file server to serve Graphiql in-browser editor
	// Reachable @ localhost:3000/
	// serve HTTP
	// pure API is reachable @ localhost:3000/graphql
	http.Handle("/", fs)
	http.ListenAndServe(":3001", nil)



}