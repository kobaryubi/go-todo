package main

import (
	"os"
	"net/http"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kobaryubi/go-todo/graph"
)

func main() {
	port := "8080"

	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:example@db:5432/postgres",
	)
	if err != nil {
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.New(conn)))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.ListenAndServe(":"+port, nil)
}
