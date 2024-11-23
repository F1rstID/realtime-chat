// domain/dto/message_dto.go
package dto

import (
	"time"

	"github.com/f1rstid/realtime-chat/domain/models"
)

// MessageResponse is a DTO for message responses
type MessageResponse struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chatId"`
	SenderID  int       `json:"senderId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewMessageResponse creates a new MessageResponse from a Message model
func NewMessageResponse(message *models.Message) *MessageResponse {
	return &MessageResponse{
		ID:        message.ID,
		ChatID:    message.ChatId,
		SenderID:  message.SenderId,
		Content:   message.Content,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}

// NewMessageResponseList creates a list of MessageResponse from Message models
func NewMessageResponseList(messages []models.Message) []MessageResponse {
	responses := make([]MessageResponse, len(messages))
	for i, message := range messages {
		responses[i] = *NewMessageResponse(&message)
	}
	return responses
}
