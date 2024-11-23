package models

import "time"

type Chat struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`

	ChatGroups []ChatGroup `json:"chatGroups" gorm:"many2many:chat_group_chats;"`
	Messages   []Message   `json:"messages" gorm:"foreignKey:chatId;"`
}
