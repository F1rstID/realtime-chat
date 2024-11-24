package repositories

import (
	"fmt"
	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ChatRepository struct {
	DB *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) repositories.ChatRepository {
	return &ChatRepository{DB: db}
}

func (r *ChatRepository) Create(chat *models.Chat) error {
	query := `INSERT INTO chats (name) VALUES ($1) RETURNING id`
	row := r.DB.QueryRow(query, chat.Name)
	return row.Scan(&chat.ID)
}

func (r *ChatRepository) FindById(id int) (*models.Chat, error) {
	chat := models.Chat{}
	query := `SELECT * FROM chats WHERE id = $1`
	err := r.DB.Get(&chat, query, id)

	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepository) Update(chat *models.Chat) error {
	query := `UPDATE chats SET name = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, chat.Name, chat.ID)
	return err
}

func (r *ChatRepository) Delete(id int) error {
	query := `DELETE FROM chats WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *ChatRepository) AddUserToChat(chatID, userID int) error {
	query := `INSERT INTO chat_groups (chatId, userId) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, chatID, userID)
	return err
}

func (r *ChatRepository) RemoveUserFromChat(chatID, userID int) error {
	query := `DELETE FROM chat_groups WHERE chatId = $1 AND userId = $2`
	_, err := r.DB.Exec(query, chatID, userID)
	return err
}

func (r *ChatRepository) GetChatUsers(chatID int) ([]models.User, error) {
	var users []models.User
	query := `
        SELECT u.* 
        FROM users u
        JOIN chat_groups cg ON u.id = cg.userId
        WHERE cg.chatId = $1
        ORDER BY u.nickname
    `
	err := r.DB.Select(&users, query, chatID)
	return users, err
}

func (r *ChatRepository) GetUserChats(userID int) ([]models.Chat, error) {
	var chats []models.Chat
	query := `
        SELECT c.* 
        FROM chats c
        JOIN chat_groups cg ON c.id = cg.chatId
        WHERE cg.userId = $1
        ORDER BY c.createdAt DESC
    `
	err := r.DB.Select(&chats, query, userID)
	return chats, err
}

func (r *ChatRepository) GetLastMessages(chatIDs []int) (map[int]*models.Message, error) {
	if len(chatIDs) == 0 {
		return make(map[int]*models.Message), nil
	}

	// Create placeholders for the IN clause
	placeholders := make([]string, len(chatIDs))
	args := make([]interface{}, len(chatIDs))
	for i := range chatIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = chatIDs[i]
	}

	query := fmt.Sprintf(`
        SELECT m.*, u.nickname as senderNickname
        FROM messages m
        JOIN users u ON m.senderId = u.id
        JOIN (
            SELECT chatId, MAX(createdAt) as maxCreatedAt
            FROM messages
            WHERE chatId IN (%s)
            GROUP BY chatId
        ) latest ON m.chatId = latest.chatId AND m.createdAt = latest.maxCreatedAt
        ORDER BY m.createdAt DESC
    `, strings.Join(placeholders, ","))

	var messages []models.Message
	err := r.DB.Select(&messages, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get last messages: %v", err)
	}

	result := make(map[int]*models.Message)
	for i := range messages {
		result[messages[i].ChatId] = &messages[i]
	}

	return result, nil
}
