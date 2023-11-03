package graph

import (
	"context"
	"os"
	"github.com/jackc/pgx/v5"
) 

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	conn *pgx.Conn
}

func New() Config {
	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:example@db:5432/postgres",
	)
	if err != nil {
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	return Config{
		Resolvers: &Resolver{conn},
	}
}
