package main

import (

	"net/http"
	"github.com/graphql-go/graphql"
	graph "github.com/boxie/wahlblock/graphql"

	"github.com/boxie/wahlblock/config"
	"github.com/graphql-go/handler"
	up "github.com/ufoscout/go-up"
)



func main() {

	cfg, _ := up.NewGoUp().
		AddFile("./config.properties", true).
		Build()


	port := cfg.GetStringOrDefault("wahlblock.port","3000")

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

	myhandler := accessControl(graphql)

	http.Handle("/graphql", myhandler)

	http.ListenAndServe(":" + port , nil)

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