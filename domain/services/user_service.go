package services

import (
	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
)

type UserService interface {
	ValidateUser(user *models.User) error
	FormatUserResponse(user *models.User) interface{}
	FilterSensitiveData(users []models.User) []interface{}
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// ValidateUser validates user data
func (s *userService) ValidateUser(user *models.User) error {
	return user.Validate()
}

// FormatUserResponse formats a single user for response
func (s *userService) FormatUserResponse(user *models.User) interface{} {
	return map[string]interface{}{
		"id":        user.ID,
		"email":     user.Email,
		"nickname":  user.Nickname,
		"createdAt": user.CreatedAt,
	}
}

// FilterSensitiveData removes sensitive information from user list
func (s *userService) FilterSensitiveData(users []models.User) []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = s.FormatUserResponse(&user)
	}
	return result
}
