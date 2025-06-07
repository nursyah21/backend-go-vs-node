package middleware

import (
	"api-go/lib"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing x-token",
		})
	}

	user, err := lib.VerifyJwt(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid token: " + err.Error(),
		})
	}

	c.Locals("user", user)

	return c.Next()
}
