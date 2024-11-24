package repositories

import (
	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	DB *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) repositories.MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) Create(message *models.Message) error {
	query := `
		INSERT INTO messages (chatId, senderId, content, createdAt, updatedAt)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	row := r.DB.QueryRow(
		query,
		message.ChatId,
		message.SenderId,
		message.Content,
		message.CreatedAt,
		message.UpdatedAt,
	)
	err := row.Scan(&message.ID)
	if err != nil {
		return err
	}

	// Fetch sender nickname
	query = `SELECT nickname FROM users WHERE id = $1`
	err = r.DB.Get(&message.SenderNickname, query, message.SenderId)
	return err
}

func (r *MessageRepository) FindById(id int) (*models.Message, error) {
	message := models.Message{}
	query := `
		SELECT m.*, u.nickname as senderNickname, m.id as id
		FROM messages m
		JOIN users u ON m.senderId = u.id
		WHERE m.id = $1
	`
	err := r.DB.Get(&message, query, id)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MessageRepository) Update(message *models.Message) error {
	// Update message
	query := `
		UPDATE messages 
		SET content = $1, updatedAt = $2
		WHERE id = $3
	`
	_, err := r.DB.Exec(query, message.Content, message.UpdatedAt, message.ID)
	if err != nil {
		return err
	}

	// Fetch sender nickname
	query = `SELECT nickname FROM users WHERE id = $1`
	return r.DB.Get(&message.SenderNickname, query, message.SenderId)
}

func (r *MessageRepository) Delete(id int) error {
	query := `DELETE FROM messages WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *MessageRepository) FindByChatId(chatId int, cursor int, limit int) ([]models.Message, error) {
	var messages []models.Message
	var query string
	var err error

	if cursor == 0 {
		// First page: get the most recent messages
		query = `
			SELECT m.*, u.nickname as senderNickname, m.id as id 
			FROM messages m
			JOIN users u ON m.senderId = u.id
			WHERE m.chatId = $1
			ORDER BY m.id DESC
			LIMIT $2
		`
		err = r.DB.Select(&messages, query, chatId, limit)
	} else {
		// Subsequent pages: get messages before the cursor
		query = `
			SELECT m.*, u.nickname as senderNickname, m.id as id 
			FROM messages m
			JOIN users u ON m.senderId = u.id
			WHERE m.chatId = $1 AND m.id < $2
			ORDER BY m.id DESC
			LIMIT $3
		`
		err = r.DB.Select(&messages, query, chatId, cursor, limit)
	}

	return messages, err
}

func (r *MessageRepository) GetLastMessageId(chatId int) (int, error) {
	var lastId int
	query := `SELECT COALESCE(MAX(id), 0) FROM messages WHERE chatId = $1`
	err := r.DB.Get(&lastId, query, chatId)
	return lastId, err
}
