package dto

import (
	"github.com/f1rstid/realtime-chat/domain/models"
	"time"
)

// LastMessageInfo represents the last message in a chat
type LastMessageInfo struct {
	MessageID      int       `json:"messageId"` // Added MessageID field
	Content        string    `json:"content"`
	SenderID       int       `json:"senderId"`
	SenderNickname string    `json:"senderNickname"`
	CreatedAt      time.Time `json:"createdAt"`
}

func NewLastMessageInfo(message *models.Message) *LastMessageInfo {
	if message == nil {
		return nil
	}
	return &LastMessageInfo{
		MessageID:      message.ID,
		Content:        message.Content,
		SenderID:       message.SenderId,
		SenderNickname: message.SenderNickname,
		CreatedAt:      message.CreatedAt,
	}
}

type UserInfo struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type ChatResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChatListResponse struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	CreatedAt   time.Time        `json:"createdAt"`
	LastMessage *LastMessageInfo `json:"lastMessage,omitempty"`
	Users       []UserInfo       `json:"users"`
}

func NewChatResponse(chat *models.Chat) *ChatResponse {
	return &ChatResponse{
		ID:        chat.ID,
		Name:      chat.Name,
		CreatedAt: chat.CreatedAt,
	}
}

func NewChatListResponse(chats []models.Chat, lastMessages map[int]*models.Message, usersMap map[int][]models.User) []ChatListResponse {
	responses := make([]ChatListResponse, len(chats))
	for i, chat := range chats {
		response := ChatListResponse{
			ID:        chat.ID,
			Name:      chat.Name,
			CreatedAt: chat.CreatedAt,
			Users:     make([]UserInfo, 0),
		}

		// Add users if available
		if users, ok := usersMap[chat.ID]; ok {
			response.Users = make([]UserInfo, len(users))
			for j, user := range users {
				response.Users[j] = UserInfo{
					ID:       user.ID,
					Nickname: user.Nickname,
				}
			}
		}

		// Add last message if available
		if lastMessage, ok := lastMessages[chat.ID]; ok {
			response.LastMessage = &LastMessageInfo{
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
