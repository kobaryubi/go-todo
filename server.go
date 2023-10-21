package main

import (
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kobaryubi/go-todo/graph"
)

func main() {
	port := "8080"

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)

	http.ListenAndServe(":"+port, nil)
}
