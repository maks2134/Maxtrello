package repository

import (
	"database/sql"
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
