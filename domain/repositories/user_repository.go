package repositories

import "github.com/f1rstid/realtime-chat/domain/models"

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindByID(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByNickname(nickname string) (*models.User, error)
	Update(user *models.User) error
	Delete(id int) error
	FindAllExcept(excludeUserId int) ([]models.User, error)
}
