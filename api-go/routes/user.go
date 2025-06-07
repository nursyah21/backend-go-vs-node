package routes

import (
	"api-go/db"

	"github.com/gofiber/fiber/v2"
)

type UserRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
}

func UserRoute(app *fiber.App, _db *db.Queries) {

	app.Post("/api/v1/music", func(c *fiber.Ctx) error {

		var req MusicRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if req.Title == "" || req.Artist == "" || req.Link == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Missing required fields: title, artist, or link",
			})
		}

		_db.CreateMusic(c.Context(), db.CreateMusicParams{
			Title: req.Title, Artist: req.Artist, Link: req.Link,
		})

		return c.JSON(fiber.Map{
			"status": "success",
		})
	})

}
