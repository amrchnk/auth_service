package repository

import (
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type User interface{
	GetUserById(id int64) (models.User, error)
	DeleteUserById(id int64) (string, error)
	GetAllUsers() ([]models.User, error)
}

type Repository struct {
	Authorization
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:NewUserPostgres(db),
	}
}
