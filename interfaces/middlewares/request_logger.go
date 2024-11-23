package middlewares

import (
	"time"

	"github.com/f1rstid/realtime-chat/infrastructure/logger"
	"github.com/gofiber/fiber/v2"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 요청 시작 시간 기록
		start := time.Now()

		// 다음 핸들러 실행
		err := c.Next()

		// 요청 처리 시간 계산
		duration := time.Since(start)

		// 요청 정보 로깅
		logger.LogRequest(
			c.Method(),
			c.Path(),
			c.IP(),
			c.Get("User-Agent"),
			c.Response().StatusCode(),
			duration,
		)

		// 에러가 발생한 경우 에러 로깅
		if err != nil {
			logger.LogError(err, c.Method(), c.Path())
		}

		return err
	}
}
