package main

import (
	"os"
	"net/http"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kobaryubi/go-todo/graph"
)

func main() {
	port := "8080"

	pool, err := pgxpool.New(
		context.Background(),
		"postgres://postgres:example@db:5432/postgres",
	)
	if err != nil {
		os.Exit(1)
	}

	defer pool.Close()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.New(pool)))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.ListenAndServe(":"+port, nil)
}
