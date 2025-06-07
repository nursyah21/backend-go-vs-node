package main

import (
	"api-go/db"
	"api-go/lib"
	"api-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	pool := lib.NewPool()
	defer pool.Close()

	db := db.New(pool)

	app := fiber.New()

	routes.HealthRoute(app)
	routes.MusicRoute(app, db)
	routes.UserRoute(app, db)

	app.Listen(":5000")
}
