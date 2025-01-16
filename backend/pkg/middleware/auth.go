package middleware

import (
	"todolist-grpc/pkg/jtoken"

	"github.com/gofiber/fiber/v2"
)

func JWTAuth() fiber.Handler {
	return JWT(jtoken.AccessTokenType)
}

func JWTRefresh() fiber.Handler {
	return JWT(jtoken.RefreshTokenType)
}

func JWT(tokenType string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(nil)
		}

		payload, err := jtoken.ValidateToken(token)
		if err != nil || payload == nil || payload["type"] != tokenType {
			return c.Status(fiber.StatusUnauthorized).JSON(nil)
		}
		c.Locals("userId", payload["id"])
		c.Locals("role", payload["role"])
		return c.Next()
	}
}
