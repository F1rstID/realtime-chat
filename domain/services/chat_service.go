package services

import (
	"encoding/json"
	"time"

	"github.com/f1rstid/realtime-chat/domain/models"
)

type ChatMessage struct {
	Type      string    `json:"type"`
	ChatID    int       `json:"chat_id"`
	SenderID  int       `json:"sender_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type ChatService struct {
	// You can add dependencies here if needed
}

func NewChatService() *ChatService {
	return &ChatService{}
}

// ParseMessage parses a raw message into a ChatMessage struct
func (cs *ChatService) ParseMessage(data []byte) (*ChatMessage, error) {
	var message ChatMessage
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, err
	}
	return &message, nil
}

// FormatMessage formats a Message model into a ChatMessage
func (cs *ChatService) FormatMessage(message *models.Message) (*ChatMessage, error) {
	chatMessage := &ChatMessage{
		Type:      "message",
		ChatID:    message.ChatId,
		SenderID:  message.SenderId,
		Content:   message.Content,
		Timestamp: message.CreatedAt,
	}

	return chatMessage, nil
}
