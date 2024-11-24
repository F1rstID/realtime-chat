package repositories

import "github.com/f1rstid/realtime-chat/domain/models"

type ChatRepository interface {
	Create(chat *models.Chat) error
	FindById(id int) (*models.Chat, error)
	Update(chat *models.Chat) error
	Delete(id int) error

	AddUserToChat(chatID, userID int) error
	RemoveUserFromChat(chatID, userID int) error
	GetChatUsers(chatID int) ([]models.User, error)
	GetUserChats(userID int) ([]models.Chat, error)
	GetLastMessages(chatIDs []int) (map[int]*models.Message, error)
}
