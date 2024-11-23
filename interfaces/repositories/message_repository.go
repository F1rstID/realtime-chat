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
	return row.Scan(&message.ID)
}

func (r *MessageRepository) FindById(id int) (*models.Message, error) {
	message := models.Message{}
	query := `SELECT * FROM messages WHERE id = $1`
	err := r.DB.Get(&message, query, id)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MessageRepository) Update(message *models.Message) error {
	query := `
		UPDATE messages 
		SET content = $1, updatedAt = $2
		WHERE id = $3
	`
	_, err := r.DB.Exec(query, message.Content, message.UpdatedAt, message.ID)
	return err
}

func (r *MessageRepository) Delete(id int) error {
	query := `DELETE FROM messages WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
