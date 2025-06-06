package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	r := fiber.New()

	r.Get("/api/v1/data", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"title":  "ayafuya",
			"artist": "rokudenashi",
			"link":   "https://yt.nurs.my.id/ayafuya",
		})
	})

	r.Listen(":5000")
}
