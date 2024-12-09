package dto

import (
	"github.com/f1rstid/realtime-chat/domain/models"
	"time"
)

type UserInfo struct {
	UserID   int    `json:"userId"` // Changed from id to userId for consistency
	Nickname string `json:"nickname"`
}

type ChatResponse struct {
	ChatID    int       `json:"chatId"` // Changed from id to chatId
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChatListResponse struct {
	ChatID      int              `json:"chatId"` // Changed from id to chatId
	Name        string           `json:"name"`
	CreatedAt   time.Time        `json:"createdAt"`
	LastMessage *LastMessageInfo `json:"lastMessage,omitempty"`
	Users       []UserInfo       `json:"users"`
}

type LastMessageInfo struct {
	MessageID      int       `json:"messageId"`
	Content        string    `json:"content"`
	SenderID       int       `json:"senderId"`
	SenderNickname string    `json:"senderNickname"`
	CreatedAt      time.Time `json:"createdAt"`
}

func NewChatResponse(chat *models.Chat) *ChatResponse {
	return &ChatResponse{
		ChatID:    chat.ID,
		Name:      chat.Name,
		CreatedAt: chat.CreatedAt,
	}
}

func NewChatListResponse(chats []models.Chat, lastMessages map[int]*models.Message, usersMap map[int][]models.User) []ChatListResponse {
	responses := make([]ChatListResponse, len(chats))
	for i, chat := range chats {
		response := ChatListResponse{
			ChatID:    chat.ID,
			Name:      chat.Name,
			CreatedAt: chat.CreatedAt,
			Users:     make([]UserInfo, 0),
		}

		// Add users if available
		if users, ok := usersMap[chat.ID]; ok {
			response.Users = make([]UserInfo, len(users))
			for j, user := range users {
				response.Users[j] = UserInfo{
					UserID:   user.ID,
					Nickname: user.Nickname,
				}
			}
		}

		// Add last message if available
		if lastMessage, ok := lastMessages[chat.ID]; ok {
			response.LastMessage = &LastMessageInfo{
				MessageID:      lastMessage.ID,
				Content:        lastMessage.Content,
				SenderID:       lastMessage.SenderId,
				SenderNickname: lastMessage.SenderNickname,
				CreatedAt:      lastMessage.CreatedAt,
			}
		}

		responses[i] = response
	}
	return responses
}
