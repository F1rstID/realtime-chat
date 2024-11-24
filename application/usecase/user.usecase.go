package usecase

import (
	"github.com/f1rstid/realtime-chat/common"
	"github.com/f1rstid/realtime-chat/domain/repositories"
	"github.com/f1rstid/realtime-chat/domain/services"
)

type UserUseCase struct {
	userRepo    repositories.UserRepository
	userService services.UserService
}

func NewUserUseCase(
	userRepo repositories.UserRepository,
	userService services.UserService,
) *UserUseCase {
	return &UserUseCase{
		userRepo:    userRepo,
		userService: userService,
	}
}

// GetAllUsersExcept retrieves all users except the specified user ID
func (uc *UserUseCase) GetAllUsersExcept(excludeUserId int) ([]common.UserListData, error) {
	// Get all users from repository
	users, err := uc.userRepo.FindAllExcept(excludeUserId)
	if err != nil {
		return nil, err
	}

	// Convert to response format
	userList := make([]common.UserListData, len(users))
	for i, user := range users {
		userList[i] = common.UserListData{
			ID:        user.ID,
			Email:     user.Email,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	return userList, nil
}
