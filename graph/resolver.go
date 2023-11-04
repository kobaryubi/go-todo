package graph

import (
	"github.com/jackc/pgx/v5"
) 

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	conn *pgx.Conn
}

func New(conn *pgx.Conn) Config {
	return Config{
		Resolvers: &Resolver{conn},
	}
}
