package dto

import (
	"time"

	"github.com/f1rstid/realtime-chat/domain/models"
)

// MessageResponse is a DTO for message responses
type MessageResponse struct {
	MessageID      int       `json:"messageId"` // Changed from "id" to "messageId"
	ChatID         int       `json:"chatId"`
	SenderID       int       `json:"senderId"`
	SenderNickname string    `json:"senderNickname"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// ChatMessagesResponse represents the response for chat messages with pagination
type ChatMessagesResponse struct {
	ChatId        int               `json:"chatId"`
	Messages      []MessageResponse `json:"messages"`
	LastMessageId int               `json:"lastMessageId"`
	HasMore       bool              `json:"hasMore"`
	NextCursor    int               `json:"nextCursor"`
}

// NewMessageResponse creates a new MessageResponse from a Message model
func NewMessageResponse(message *models.Message) *MessageResponse {
	return &MessageResponse{
		MessageID:      message.ID, // Changed from ID to MessageID
		ChatID:         message.ChatId,
		SenderID:       message.SenderId,
		SenderNickname: message.SenderNickname,
		Content:        message.Content,
		CreatedAt:      message.CreatedAt,
		UpdatedAt:      message.UpdatedAt,
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
