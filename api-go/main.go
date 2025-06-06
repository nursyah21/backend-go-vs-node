package main

import (
	"api-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.HealthRoute(app)
	routes.MusicRoute(app)

	app.Listen(":5000")
}
