package lib

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool() *pgxpool.Pool {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig("user=go_user password=go_password port=5433 dbname=api_go sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Failed to create pool: %v", err)
	}

	return pool
}
