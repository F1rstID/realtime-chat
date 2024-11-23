// interfaces/controllers/chat_controller.go
package controllers

import (
	"github.com/f1rstid/realtime-chat/application/usecase"
	"github.com/f1rstid/realtime-chat/interfaces"
	"github.com/gofiber/fiber/v2"
)

type CreatePrivateChatRequest struct {
	TargetId int `json:"targetId" example:"1"`
}

// CreateGroupChatRequest represents the request for creating a group chat
type CreateGroupChatRequest struct {
	Name    string `json:"name" example:"Team Chat" validate:"required"`
	UserIDs []int  `json:"userIds" example:"[1,2,3] validate:"required"`
}

type ChatController struct {
	chatUseCase    *usecase.ChatUsecase
	messageUseCase *usecase.MessageUsecase
}

func NewChatController(
	chatUseCase *usecase.ChatUsecase,
	messageUseCase *usecase.MessageUsecase,
) *ChatController {
	return &ChatController{
		chatUseCase:    chatUseCase,
		messageUseCase: messageUseCase,
	}
}

// CreatePrivateChat godoc
// @Summary      1:1 채팅 생성
// @Description  두 사용자 간의 1:1 채팅을 생성합니다
// @Tags         Chat
// @Accept       json
// @Produce      json
// @Param        user1_id path int true "사용자1 ID"
// @Param        user2_id path int true "사용자2 ID"
// @Success      201  {object}  common.ChatResponse
// @Failure      400  {object}  common.ErrInvalidRequest
// @Failure      500  {object}  common.ErrInternalServer
// @Security     Bearer
// @Router       /api/chats/private/{user1_id}/{user2_id} [post]
func (cc *ChatController) CreatePrivateChat(c *fiber.Ctx) error {
	var req CreatePrivateChatRequest
	if err := c.BodyParser(&req); err != nil {
		return interfaces.SendBadRequest(c, "잘못된 요청 형식입니다")
	}

	userID := c.Locals("userId").(int)

	chat, err := cc.chatUseCase.CreatePrivateChat(userID, req.TargetId)
	if err != nil {
		switch err.Error() {
		case "user1 not found", "user2 not found":
			return interfaces.SendNotFound(c, "사용자")
		default:
			return interfaces.SendInternalError(c)
		}
	}

	return interfaces.SendCreated(c, chat)
}

// CreateGroupChat godoc
// @Summary      그룹 채팅 생성
// @Description  그룹 채팅방을 생성합니다
// @Tags         Chat
// @Accept       json
// @Produce      json
// @Param        request body CreateGroupChatRequest true "채팅방 생성 정보"
// @Success      201  {object}  common.ChatResponse
// @Failure      400  {object}  common.ErrInvalidRequest
// @Failure      500  {object}  common.ErrInternalServer
// @Security     Bearer
// @Router       /api/chats/group [post]
func (cc *ChatController) CreateGroupChat(c *fiber.Ctx) error {
	var req CreateGroupChatRequest
	if err := c.BodyParser(&req); err != nil {
		return interfaces.SendBadRequest(c, "잘못된 요청 형식입니다")
	}

	if req.Name == "" {
		return interfaces.SendBadRequest(c, "채팅방 이름은 필수 항목입니다")
	}

	if len(req.UserIDs) < 2 {
		return interfaces.SendBadRequest(c, "그룹 채팅은 최소 2명 이상의 사용자가 필요합니다")
	}

	chat, err := cc.chatUseCase.CreateGroupChat(req.Name, req.UserIDs)
	if err != nil {
		switch err.Error() {
		case "user not found":
			return interfaces.SendNotFound(c, "사용자")
		default:
			return interfaces.SendInternalError(c)
		}
	}

	return interfaces.SendCreated(c, chat)
}
