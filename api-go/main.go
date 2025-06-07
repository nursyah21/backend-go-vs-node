package main

import (
	"api-go/db"
	"api-go/routes"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	poolConfig, err := pgxpool.ParseConfig("user=go_user password=go_password port=5433 dbname=api_go sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
		return
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
	defer pool.Close()

	db := db.New(pool)

	app := fiber.New()

	routes.HealthRoute(app)
	routes.MusicRoute(app, db)

	app.Listen(":5000")
}
