// application/usecase/auth_usecase.go

package usecase

import (
	"errors"
	"strings"
	"time"

	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
	"github.com/f1rstid/realtime-chat/domain/services"
)

// RegisterInput defines the input data for registration
type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Nickname string `json:"nickname" validate:"required,min=2,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

// LoginInput defines the input data for login
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse defines the response data for authentication operations
type AuthResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

type AuthUsecase struct {
	userRepo    repositories.UserRepository
	authService services.AuthService
}

func NewAuthUsecase(userRepo repositories.UserRepository, authService services.AuthService) *AuthUsecase {
	return &AuthUsecase{
		userRepo:    userRepo,
		authService: authService,
	}
}

func (au *AuthUsecase) Register(input RegisterInput) (*AuthResponse, error) {
	// Clean input
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))
	input.Nickname = strings.TrimSpace(input.Nickname)

	// Create user model
	user := &models.User{
		Email:     input.Email,
		Nickname:  input.Nickname,
		Password:  input.Password,
		CreatedAt: time.Now(),
	}

	// Check if email exists
	existingUser, err := au.userRepo.FindByEmail(input.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Check if nickname exists
	existingUser, err = au.userRepo.FindByNickname(input.Nickname)
	if err == nil && existingUser != nil {
		return nil, errors.New("nickname already exists")
	}

	// Hash password
	hashedPassword, err := au.authService.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	// Create user
	createdUser, err := au.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	// Generate token
	token, err := au.authService.GenerateToken(createdUser)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  createdUser,
	}, nil
}

func (au *AuthUsecase) Login(input LoginInput) (*AuthResponse, error) {
	// Clean input
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))

	// Find user by email
	user, err := au.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	err = au.authService.ComparePassword(user.Password, input.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate token
	token, err := au.authService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
