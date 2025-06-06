package main

import (
	"api-go/db"
	"api-go/routes"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=go_user password=go_password port=5433 dbname=api_go sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	defer conn.Close(ctx)

	db := db.New(conn)

	app := fiber.New()

	routes.HealthRoute(app)
	routes.MusicRoute(app, db)

	app.Listen(":5000")
}
