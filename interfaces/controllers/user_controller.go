package controllers

import (
	"github.com/f1rstid/realtime-chat/application/usecase"
	"github.com/f1rstid/realtime-chat/interfaces"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUseCase *usecase.UserUseCase
}

func NewUserController(userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

// GetAllUsers godoc
// @Summary      전체 사용자 목록 조회
// @Description  현재 로그인한 사용자를 제외한 전체 사용자 목록을 조회합니다
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  common.UserListResponse
// @Failure      401  {object}  common.ErrUnauthorized
// @Failure      500  {object}  common.ErrInternalServer
// @Security     Bearer
// @Router       /api/users [get]
func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	currentUserId := c.Locals("userId").(int)

	users, err := uc.userUseCase.GetAllUsersExcept(currentUserId)
	if err != nil {
		return interfaces.SendInternalError(c)
	}

	return interfaces.SendSuccess(c, users)
}
