package services

import (
	"errors"
	"time"

	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type TokenClaims struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

type AuthService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
	GenerateToken(user *models.User) (string, error)
	ValidateToken(tokenString string) (*TokenClaims, error)
	RefreshToken(tokenString string) (string, error)
}

type authService struct {
	jwtSecret       string
	tokenExpiration time.Duration
	refreshDuration time.Duration
	passwordCost    int
}

func NewAuthService(jwtSecret string) AuthService {
	return &authService{
		jwtSecret:       jwtSecret,
		tokenExpiration: 24 * time.Hour, // Token expires in 24 hours
		refreshDuration: 72 * time.Hour, // Refresh token valid for 72 hours
		passwordCost:    5,              // Higher cost = more secure but slower
	}
}

func (a *authService) HashPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), a.passwordCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedBytes), nil
}

func (a *authService) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (a *authService) GenerateToken(user *models.User) (string, error) {
	claims := TokenClaims{
		UserID:   user.ID,
		Email:    user.Email,
		Nickname: user.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.tokenExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return signedToken, nil
}

func (a *authService) ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (a *authService) RefreshToken(tokenString string) (string, error) {
	claims, err := a.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Check if token is eligible for refresh
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > -a.refreshDuration {
		// Create new token
		user := &models.User{
			ID:       claims.UserID,
			Email:    claims.Email,
			Nickname: claims.Nickname,
		}
		return a.GenerateToken(user)
	}

	return "", errors.New("refresh token expired")
}
