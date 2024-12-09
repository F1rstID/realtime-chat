package controllers

import (
	"github.com/f1rstid/realtime-chat/application/usecase"
	"github.com/f1rstid/realtime-chat/interfaces"
	"github.com/gofiber/fiber/v2"
)

type CreatePrivateChatRequest struct {
	TargetId int `json:"targetId" example:"1"`
}

// @Description 그룹 채팅방 생성 요청
type CreateGroupChatRequest struct {
	// 채팅방 이름
	Name string `json:"name" example:"Team Chat"`
	// 초대할 사용자 ID 목록
	UserIDs []int `json:"userIds" example:"1,2,3"`
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

// GetUserChats godoc
// @Summary      사용자의 채팅방 목록 조회
// @Description  현재 로그인한 사용자가 참여중인 모든 채팅방 목록을 조회합니다. 각 채팅방의 마지막 메시지 정보와 참여자 정보도 함께 제공됩니다.
// @Tags         Chat
// @Accept       json
// @Produce      json
// @Success      200  {object}  common.ChatListResponse
// @Failure      401  {object}  common.ErrUnauthorized
// @Failure      500  {object}  common.ErrInternalServer
// @Security     Bearer
// @Router       /api/chats [get]
func (cc *ChatController) GetUserChats(c *fiber.Ctx) error {
	userID := c.Locals("userId").(int)

	chats, err := cc.chatUseCase.GetUserChats(userID)
	if err != nil {
		return interfaces.SendInternalError(c)
	}

	return interfaces.SendSuccess(c, chats)
}

// CreatePrivateChat godoc
// @Summary      1:1 채팅 생성
// @Description  두 사용자 간의 1:1 채팅을 생성합니다
// @Tags         Chat
// @Accept       json
// @Produce      json
// @Param        request body CreatePrivateChatRequest true "상대 사용자ID"
// @Success      201  {object}  common.ChatResponse
// @Failure      400  {object}  common.ErrInvalidRequest
// @Failure      500  {object}  common.ErrInternalServer
// @Security     Bearer
// @Router       /api/chats/private [post]
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
// @Param        request body CreateGroupChatRequest true "채팅방 이름, 참여자 ID 목록"
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

	if len(req.UserIDs) < 1 {
		return interfaces.SendBadRequest(c, "초대할 사용자가 한 명 이상 필요합니다")
	}

	// Get current user ID from context
	currentUserID := c.Locals("userId").(int)

	// Add current user to the userIDs if not already included
	hasCurrentUser := false
	for _, id := range req.UserIDs {
		if id == currentUserID {
			hasCurrentUser = true
			break
		}
	}

	if !hasCurrentUser {
		req.UserIDs = append(req.UserIDs, currentUserID)
	}

	chat, err := cc.chatUseCase.CreateGroupChat(req.Name, req.UserIDs)
	if err != nil {
		switch err.Error() {
		case "user not found":
			return interfaces.SendNotFound(c, "사용자")
		case "chat name is required":
			return interfaces.SendBadRequest(c, "채팅방 이름은 필수 항목입니다")
		case "at least one other user is required":
			return interfaces.SendBadRequest(c, "초대할 사용자가 한 명 이상 필요합니다")
		default:
			return interfaces.SendInternalError(c)
		}
	}

	return interfaces.SendCreated(c, chat)
}

func (cc *ChatController) GetChats(c *fiber.Ctx) error {
	userID := c.Locals("userId").(int)

	chats, err := cc.chatUseCase.GetUserChats(userID)
	if err != nil {
		return interfaces.SendInternalError(c)
	}

	return interfaces.SendSuccess(c, chats)
}
