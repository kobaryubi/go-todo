package graph

import (
	"github.com/jackc/pgx/v5/pgxpool"
) 

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) Config {
	return Config{
		Resolvers: &Resolver{pool},
	}
}
