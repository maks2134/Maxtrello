package repository

import (
	"database/sql"
	"errors"
	"user-service/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3) 
	          RETURNING id`

	return r.db.QueryRow(query, user.Email, user.Password, user.Name).Scan(
		&user.ID,
	)
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password, name
	          FROM users WHERE email = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return user, err
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
	query := `SELECT id, email, password, name
	          FROM users WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return user, err
}
