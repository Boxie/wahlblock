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

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin","*")
	})

	http.Handle("/graphql", graphql)

	http.ListenAndServe(":3000", mux)

}