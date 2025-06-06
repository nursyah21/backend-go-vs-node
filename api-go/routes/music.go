package routes

import "github.com/gofiber/fiber/v2"

func MusicRoute(app *fiber.App) {

	app.Get("/api/v1/music", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status": "ok",
		})
	})

}
