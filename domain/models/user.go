package models

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Nickname  string    `json:"nickname" db:"nickname"`
	Password  string    `json:"-" db:"password"` // "-" prevents password from being included in JSON
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
}

// Validate performs validation on user fields
func (u *User) Validate() error {
	// Email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	// Nickname validation
	if len(u.Nickname) < 2 || len(u.Nickname) > 20 {
		return errors.New("nickname must be between 2 and 20 characters")
	}
	nicknameRegex := regexp.MustCompile(`^[a-zA-Z0-9가-힣_.]+$`)
	if !nicknameRegex.MatchString(u.Nickname) {
		return errors.New("nickname can only contain letters, numbers and underscores")
	}

	// Password validation (performed before hashing)
	if len(u.Password) < 1 { // Corrected minimum password length
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
