package repositories

import "github.com/f1rstid/realtime-chat/domain/models"

type MessageRepository interface {
	Create(message *models.Message) error
	FindById(id int) (*models.Message, error)
	Update(message *models.Message) error
	Delete(id int) error
	FindByChatId(chatId int, cursor int, limit int) ([]models.Message, error)
	GetLastMessageId(chatId int) (int, error)
}
