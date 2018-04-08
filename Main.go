package main

import (

	"net/http"
	"github.com/graphql-go/graphql"
	graph "github.com/boxie/wahlblock/graphql"

	"github.com/graphql-go/handler"
)



func main() {

	//Assign graphql to http handler
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: graph.RootQuery,
		Mutation: graph.RootMutation,
	})

	graphql := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,

	})

	myhandler := accessControl(graphql)

	http.Handle("/graphql", myhandler)

	http.ListenAndServe(":3000", nil)

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, content-type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}