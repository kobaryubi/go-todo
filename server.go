package main

import (
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kobaryubi/go-todo/graph"
)

func main() {
	port := "8080"

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.New()))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.ListenAndServe(":"+port, nil)
}
