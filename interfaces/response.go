// interfaces/response.go
package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

// Response 구조체는 공통 응답 구조를 사용합니다.
type Response struct {
	Success bool        `json:"success" example:"true"`
	Code    int         `json:"code" example:"2000"`
	Data    interface{} `json:"data"`
}

// Data models are moved to the common package, so you can remove redundant definitions here.

const (
	// Success codes (2xxx)
	StatusSuccess = 2000
	StatusCreated = 2001

	// Client error codes (4xxx)
	StatusBadRequest         = 4000
	StatusUnauthorized       = 4001
	StatusForbidden          = 4002
	StatusNotFound           = 4003
	StatusEmailExists        = 4004
	StatusNicknameExists     = 4005
	StatusInvalidCredentials = 4006
	StatusValidationError    = 4007
	StatusInvalidToken       = 4008

	// Server error codes (5xxx)
	StatusInternalError = 5000
	StatusDBError       = 5001
)

// Response helpers
func SendResponse(c *fiber.Ctx, httpStatus, code int, success bool, data interface{}) error {
	return c.Status(httpStatus).JSON(Response{
		Code:    code,
		Success: success,
		Data:    data,
	})
}

func SendSuccess(c *fiber.Ctx, data interface{}) error {
	return SendResponse(c, fiber.StatusOK, StatusSuccess, true, data)
}

func SendCreated(c *fiber.Ctx, data interface{}) error {
	return SendResponse(c, fiber.StatusCreated, StatusCreated, true, data)
}

func SendError(c *fiber.Ctx, httpStatus, code int, message string) error {
	return SendResponse(c, httpStatus, code, false, message)
}

// Error helpers
func SendBadRequest(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusBadRequest, StatusBadRequest, message)
}

func SendUnauthorized(c *fiber.Ctx) error {
	return SendError(c, fiber.StatusUnauthorized, StatusUnauthorized, "인증되지 않은 접근입니다")
}

func SendForbidden(c *fiber.Ctx) error {
	return SendError(c, fiber.StatusForbidden, StatusForbidden, "접근 권한이 없습니다")
}

func SendNotFound(c *fiber.Ctx, resource string) error {
	return SendError(c, fiber.StatusNotFound, StatusNotFound, resource+"를 찾을 수 없습니다")
}

func SendEmailExists(c *fiber.Ctx) error {
	return SendError(c, fiber.StatusConflict, StatusEmailExists, "이미 사용중인 이메일입니다")
}

func SendInternalError(c *fiber.Ctx) error {
	return SendError(c, fiber.StatusInternalServerError, StatusInternalError, "내부 서버 오류가 발생했습니다")
}
