package routes

import (
	"api-go/db"
	"api-go/lib"
	"api-go/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MusicRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
}

func parseAndValidateRequest(c *fiber.Ctx) error {
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

	c.Locals("req", req)

	return c.Next()
}

func MusicRoute(app *fiber.App, _db *db.Queries) {

	app.Get("/api/v1/music", func(c *fiber.Ctx) error {
		musics, err := _db.ListMusics(c.Context())

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(musics)
	})

	app.Post("/api/v1/music", middleware.AuthMiddleware, parseAndValidateRequest,
		func(c *fiber.Ctx) error {

			user := c.Locals("user").(*lib.CustomClaims)
			req := c.Locals("req").(MusicRequest)

			userId, _ := strconv.ParseInt(user.ID, 10, 64)

			if err := _db.CreateMusic(c.Context(), db.CreateMusicParams{
				Userid: userId,
				Title:  req.Title, Artist: req.Artist, Link: req.Link,
			}); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"status": "success",
			})
		})

	app.Put("/api/v1/music/:id", middleware.AuthMiddleware, parseAndValidateRequest,
		middleware.ParamId, func(c *fiber.Ctx) error {

			user := c.Locals("user").(*lib.CustomClaims)
			req := c.Locals("req").(MusicRequest)
			id := c.Locals("id").(int64)

			userId, _ := strconv.ParseInt(user.ID, 10, 64)

			_db.UpdateMusic(c.Context(), db.UpdateMusicParams{
				ID: id, Userid: userId, Title: req.Title, Artist: req.Artist, Link: req.Link,
			})

			return c.JSON(fiber.Map{
				"status": "success",
			})
		})

	app.Delete("/api/v1/music/:id", middleware.AuthMiddleware, middleware.ParamId,
		func(c *fiber.Ctx) error {
			user := c.Locals("user").(*lib.CustomClaims)
			id := c.Locals("id").(int64)

			userId, _ := strconv.ParseInt(user.ID, 10, 64)

			_db.DeleteMusic(c.Context(), db.DeleteMusicParams{
				ID: id, Userid: userId,
			})

			return c.JSON(fiber.Map{
				"status": "success",
			})
		})
}
