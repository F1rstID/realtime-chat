// interfaces/controllers/auth_controller.go
package controllers

import (
	"github.com/f1rstid/realtime-chat/application/usecase"
	"github.com/f1rstid/realtime-chat/common"
	"github.com/gofiber/fiber/v2"
)

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
	Email    string `json:"email" example:"user@example.com" validate:"required,email"`
	Nickname string `json:"nickname" example:"홍길동" validate:"required,min=2,max=20"`
	Password string `json:"password" example:"password123" validate:"required,min=8"`
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com" validate:"required,email"`
	Password string `json:"password" example:"password123" validate:"required"`
}

type AuthController struct {
	authUseCase *usecase.AuthUsecase
}

func NewAuthController(usecase *usecase.AuthUsecase) *AuthController {
	return &AuthController{
		authUseCase: usecase,
	}
}

// Register godoc
// @Summary      사용자 등록
// @Description  새로운 사용자를 등록하고 인증 토큰을 반환합니다
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body RegisterRequest true "등록 정보"
// @Success      201  {object}  common.RegisterResponse
// @Failure      400  {object}  common.ErrInvalidRequest
// @Failure      409  {object}  common.ErrEmailExists
// @Failure      409  {object}  common.ErrNicknameExists
// @Failure      500  {object}  common.ErrInternalServer
// @Router       /api/auth/register [post]
func (ac *AuthController) Register(c *fiber.Ctx) error {
	var input RegisterRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.InvalidRequest)
	}

	if input.Email == "" || input.Password == "" || input.Nickname == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.InvalidRequest)
	}

	registerInput := usecase.RegisterInput{
		Email:    input.Email,
		Password: input.Password,
		Nickname: input.Nickname,
	}

	authResponse, err := ac.authUseCase.Register(registerInput)
	if err != nil {
		switch err.Error() {
		case "email already exists":
			return c.Status(fiber.StatusConflict).JSON(common.EmailExists)
		case "nickname already exists":
			return c.Status(fiber.StatusConflict).JSON(common.NicknameExists)
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(common.InternalServer)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(common.BaseResponse{
		Success: true,
		Code:    common.StatusCreated,
		Data:    authResponse,
	})
}

// Login godoc
// @Summary      사용자 로그인
// @Description  이메일과 비밀번호로 로그인하고 인증 토큰을 반환합니다
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "로그인 정보"
// @Success      200  {object}  common.LoginResponse
// @Failure      400  {object}  common.ErrInvalidRequest
// @Failure      401  {object}  common.ErrInvalidAuth
// @Failure      500  {object}  common.ErrInternalServer
// @Router       /api/auth/login [post]
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var input LoginRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.InvalidRequest)
	}

	if input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.InvalidRequest)
	}

	loginInput := usecase.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	}

	authResponse, err := ac.authUseCase.Login(loginInput)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(common.InvalidAuth)
	}

	return c.Status(fiber.StatusOK).JSON(common.BaseResponse{
		Success: true,
		Code:    common.StatusSuccess,
		Data:    authResponse,
	})
}
