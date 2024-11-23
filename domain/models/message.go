package models

import "time"

type Message struct {
	ID             int       `json:"id" db:"id"`
	ChatId         int       `json:"chatId" db:"chatId"`
	SenderId       int       `json:"senderId" db:"senderId"`
	SenderNickname string    `json:"senderNickname" db:"senderNickname"`
	Content        string    `json:"content" db:"content"`
	CreatedAt      time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updatedAt"`

	Chat   Chat `json:"chat" gorm:"foreignKey:chatId;"`
	Sender User `json:"sender" gorm:"foreignKey:senderId;"`
}
