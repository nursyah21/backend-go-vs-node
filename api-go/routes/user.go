package routes

import (
	"api-go/db"
	"api-go/lib"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserRoute(app *fiber.App, _db *db.Queries) {

	app.Post("/api/v1/register", func(c *fiber.Ctx) error {
		var req UserRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if req.Username == "" || req.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Missing required fields: username, password",
			})
		}

		password, _ := lib.HashPassword(req.Password)

		err := _db.CreateUser(c.Context(), db.CreateUserParams{
			Username: req.Username, Password: password,
		})

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": req.Username + " already exist",
			})
		}

		return c.JSON(fiber.Map{
			"status": "success",
		})
	})

	app.Post("/api/v1/login", func(c *fiber.Ctx) error {
		var req UserRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if req.Username == "" || req.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Missing required fields: username, password",
			})
		}

		res, _ := _db.GetUser(c.Context(), req.Username)

		if err := lib.VerifyPassword(res.Password, req.Password); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "username and password didnt match",
			})
		}

		id := strconv.Itoa(int(res.ID))
		token, _ := lib.GenerateJwt(id, res.Username)

		return c.JSON(fiber.Map{
			"status": "success",
			"token":  token,
		})
	})

}
