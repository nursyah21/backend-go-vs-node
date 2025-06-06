package routes

import (
	"api-go/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func MusicRoute(app *fiber.App, _db *db.Queries) {

	app.Get("/api/v1/music", func(c *fiber.Ctx) error {
		musics, err := _db.ListMusics(c.Context())

		if err != nil {
			return c.JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"data": musics,
		})
	})

	app.Post("/api/v1/music", func(c *fiber.Ctx) error {
		type Request struct {
			Title  string `json:"title"`
			Artist string `json:"artist"`
			Link   string `json:"link"`
		}

		var req Request
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

		return c.JSON(&fiber.Map{
			"status": "success",
		})
	})

	app.Put("/api/v1/music/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid music ID",
			})
		}

		type Request struct {
			Title  string `json:"title"`
			Artist string `json:"artist"`
			Link   string `json:"link"`
		}

		var req Request
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

		_db.UpdateMusic(c.Context(), db.UpdateMusicParams{
			ID: id, Title: req.Title, Artist: req.Artist, Link: req.Link,
		})

		return c.JSON(&fiber.Map{
			"status": "success",
		})
	})

	app.Delete("/api/v1/music/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid music ID",
			})
		}

		_db.DeleteMusic(c.Context(), id)

		return c.JSON(&fiber.Map{
			"status": "success",
		})
	})
}
