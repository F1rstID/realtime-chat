package middlewares

import (
	"strings"

	"github.com/f1rstid/realtime-chat/domain/services"
	"github.com/f1rstid/realtime-chat/interfaces"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(authService services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from header
		authHeader := c.Get("authorization")

		if authHeader == "" {
			return interfaces.SendUnauthorized(c)
		}

		// Check Bearer format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return interfaces.SendUnauthorized(c)
		}

		// Validate token
		claims, err := authService.ValidateToken(parts[1])
		if err != nil {
			return interfaces.SendUnauthorized(c)
		}

		// Set claims in context
		c.Locals("userId", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userNickname", claims.Nickname)

		return c.Next()
	}
}
