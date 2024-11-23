package middlewares

import (
	"github.com/f1rstid/realtime-chat/infrastructure/logger"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		// 에러 코드 설정
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		// 에러 로깅
		logger.Error("Request failed - Method: %s, Path: %s, Error: %v",
			c.Method(), c.Path(), err)

		// 클라이언트에게 에러 응답
		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
			"code":    code,
		})
	}
}
