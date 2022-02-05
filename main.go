package main

import (
	"gql/db"
	"gql/graph"
	"gql/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	psqlDb := &db.Psql{}
	psqlDb.Connect()
	defer (*psqlDb).Db.Close()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Db: psqlDb}}))

	http.Handle("/query", srv)

	log.Printf("Listening on PORT %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
