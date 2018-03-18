package main

import (

	"net/http"
	"github.com/graphql-go/graphql"
	graph "github.com/boxie/wahlblock/graphql"

	"github.com/boxie/wahlblock/config"
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


	// static file server to serve Graphiql in-browser editor
	// Reachable @ localhost:3000/
	fs := http.FileServer(http.Dir("static"))

	// serve HTTP
	// pure API is reachable @ localhost:3000/graphql
	http.Handle("/graphql", graphql)
	http.Handle("/", fs)
	http.ListenAndServe(":5000", nil)

}