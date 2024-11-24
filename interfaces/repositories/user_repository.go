package repositories

import (
	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repositories.UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (email, nickname, password) VALUES ($1, $2, $3) RETURNING id`
	row := r.DB.QueryRow(query, user.Email, user.Nickname, user.Password)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE id = $1`
	err := r.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE email = $1`
	err := r.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByNickname(nickname string) (*models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE nickname = $1`
	err := r.DB.Get(&user, query, nickname)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := `UPDATE users SET email = $1, nickname = $2, password = $3 WHERE id = $4`
	_, err := r.DB.Exec(query, user.Email, user.Nickname, user.Password, user.ID)
	return err
}

// Add the missing Delete method
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *UserRepository) FindAllExcept(excludeUserId int) ([]models.User, error) {
	var users []models.User
	query := `SELECT id, email, nickname, createdAt FROM users WHERE id != $1 ORDER BY createdAt DESC`
	err := r.DB.Select(&users, query, excludeUserId)
	return users, err
}
