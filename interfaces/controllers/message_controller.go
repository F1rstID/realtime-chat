package controllers

import (
	"github.com/f1rstid/realtime-chat/application/usecase"
	"github.com/f1rstid/realtime-chat/interfaces"
	"github.com/gofiber/fiber/v2"
)

// SendMessageRequest represents the request for sending a message
type SendMessageRequest struct {
	ChatID  int    `json:"chatId" example:"1" validate:"required"`
	Content string `json:"content" example:"Hello, how are you?" validate:"required"`
}

// UpdateMessageRequest represents the request for updating a message
type UpdateMessageRequest struct {
	Content string `json:"content" example:"Updated message content" validate:"required"`
}

type MessageController struct {
	messageUseCase *usecase.MessageUsecase
}

func NewMessageController(messageUseCase *usecase.MessageUsecase) *MessageController {
	return &MessageController{
		messageUseCase: messageUseCase,
	}
}

// SendMessage godoc
// @Summary      메시지 전송
// @Description  채팅방에 새로운 메시지를 전송합니다
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        request body SendMessageRequest true "메시지 정보"
// @Success      201  {object}  common.MessageResponse
// @Failure      400  {object}  common.ErrInvalidRequest
// @Failure      500  {object}  common.ErrInternalServer
// @Security     Bearer
// @Router       /api/messages [post]
func (mc *MessageController) SendMessage(c *fiber.Ctx) error {
	var req SendMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return interfaces.SendBadRequest(c, "잘못된 요청 형식입니다")
	}

	if req.Content == "" {
		return interfaces.SendBadRequest(c, "메시지 내용은 필수 항목입니다")
	}

	userID := c.Locals("userId").(int)

	message, err := mc.messageUseCase.SendMessage(req.ChatID, userID, req.Content)
	if err != nil {
		switch err.Error() {
		case "chat not found":
			return interfaces.SendNotFound(c, "채팅방")
		default:
			return interfaces.SendInternalError(c)
		}
	}

	return interfaces.SendCreated(c, message)
}

// UpdateMessage godoc
// @Summary      메시지 수정
// @Description  기존 메시지의 내용을 수정합니다
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "메시지 ID"
// @Param        request body UpdateMessageRequest true "수정할 메시지 내용"
// @Success      200  {object}  common.MessageResponse
// @Failure      400  {object}  common.ErrInvalidRequest
// @Failure      403  {object}  common.ErrUnauthorizedMessage
// @Failure      500  {object}  common.ErrInternalServer
// @Security     Bearer
// @Router       /api/messages/{id} [put]
func (mc *MessageController) UpdateMessage(c *fiber.Ctx) error {
	messageID, err := c.ParamsInt("id")
	if err != nil {
		return interfaces.SendBadRequest(c, "잘못된 메시지 ID입니다")
	}

	var req UpdateMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return interfaces.SendBadRequest(c, "잘못된 요청 형식입니다")
	}

	if req.Content == "" {
		return interfaces.SendBadRequest(c, "메시지 내용은 필수 항목입니다")
	}

	userID := c.Locals("userId").(int)

	message, err := mc.messageUseCase.UpdateMessage(messageID, userID, req.Content)
	if err != nil {
		switch err.Error() {
		case "message not found":
			return interfaces.SendNotFound(c, "메시지")
		case "unauthorized to update this message":
			return interfaces.SendForbidden(c)
		default:
			return interfaces.SendInternalError(c)
		}
	}

	return interfaces.SendSuccess(c, message)
}

// DeleteMessage godoc
// @Summary      메시지 삭제
// @Description  메시지를 삭제합니다
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "메시지 ID"
// @Success      200  {object}  common.MessageResponse
// @Failure      400  {object}  common.ErrInvalidRequest
// @Failure      403  {object}  common.ErrUnauthorizedMessage
// @Failure      500  {object}  common.ErrInternalServer
// @Security     Bearer
// @Router       /api/messages/{id} [delete]
func (mc *MessageController) DeleteMessage(c *fiber.Ctx) error {
	messageID, err := c.ParamsInt("id")
	if err != nil {
		return interfaces.SendBadRequest(c, "잘못된 메시지 ID입니다")
	}

	userID := c.Locals("userId").(int)

	if err := mc.messageUseCase.DeleteMessage(messageID, userID); err != nil {
		switch err.Error() {
		case "message not found":
			return interfaces.SendNotFound(c, "메시지")
		case "unauthorized to delete this message":
			return interfaces.SendForbidden(c)
		default:
			return interfaces.SendInternalError(c)
		}
	}

	return interfaces.SendSuccess(c, "메시지가 삭제되었습니다")
}
