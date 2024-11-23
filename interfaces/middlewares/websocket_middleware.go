// interfaces/middlewares/websocket_middleware.go
package middlewares

import (
	"log"
	"strings"

	"github.com/f1rstid/realtime-chat/domain/services"
	"github.com/gofiber/fiber/v2"
)

func WebSocketAuthMiddleware(authService services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Printf("WebSocket auth middleware - Path: %s", c.Path())

		// Get token from query parameter or header
		var token string

		// First try to get from query parameter
		token = c.Query("token")
		log.Printf("Token from query: %s", maskToken(token))

		// If not in query, try to get from header
		if token == "" {
			auth := c.Get("Authorization")
			log.Printf("Authorization header: %s", maskToken(auth))
			if auth != "" && strings.HasPrefix(auth, "Bearer ") {
				token = auth[7:] // Remove "Bearer " prefix
			}
		}

		if token == "" {
			log.Println("WebSocket auth failed: no token provided")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authentication required",
			})
		}

		// Validate token
		claims, err := authService.ValidateToken(token)
		if err != nil {
			log.Printf("WebSocket auth failed: invalid token - %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		log.Printf("WebSocket auth successful - UserID: %d", claims.UserID)

		// Store user information in context
		c.Locals("userId", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userNickname", claims.Nickname)

		return c.Next()
	}
}

// maskToken masks the token for logging purposes
func maskToken(token string) string {
	if len(token) > 10 {
		return token[:4] + "..." + token[len(token)-4:]
	}
	return "[empty]"
}
