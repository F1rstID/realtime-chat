// domain/dto/chat_dto.go
package dto

import (
	"github.com/f1rstid/realtime-chat/domain/models"
	"time"
)

type ChatResponse struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	CreatedAt   time.Time        `json:"createdAt"`
	LastMessage *LastMessageInfo `json:"lastMessage,omitempty"`
}

type LastMessageInfo struct {
	Content        string    `json:"content"`
	SenderID       int       `json:"senderId"`
	SenderNickname string    `json:"senderNickname"`
	SentAt         time.Time `json:"sentAt"`
}

func NewChatResponse(chat *models.Chat, lastMessage *models.Message) *ChatResponse {
	response := &ChatResponse{
		ID:        chat.ID,
		Name:      chat.Name,
		CreatedAt: chat.CreatedAt,
	}

	if lastMessage != nil {
		response.LastMessage = &LastMessageInfo{
			Content:        lastMessage.Content,
			SenderID:       lastMessage.SenderId,
			SenderNickname: lastMessage.SenderNickname,
			SentAt:         lastMessage.CreatedAt,
		}
	}

	return response
}

func NewChatListResponse(chats []models.Chat, lastMessages map[int]*models.Message) []ChatResponse {
	responses := make([]ChatResponse, len(chats))
	for i, chat := range chats {
		lastMessage := lastMessages[chat.ID]
		responses[i] = *NewChatResponse(&chat, lastMessage)
	}
	return responses
}
