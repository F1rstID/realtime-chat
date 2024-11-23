package models

import "time"

type ChatGroup struct {
	Name      string    `json:"name" db:"name"`
	UserId    int       `json:"userId" db:"userId"`
	ChatId    int       `json:"chatId" db:"chatId"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`

	Users []User `json:"users" gorm:"many2many:chat_group_users;"`
	Chats []Chat `json:"chats" gorm:"many2many:chat_group_chats;"`
}
