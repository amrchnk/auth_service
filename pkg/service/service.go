package service

import (
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/amrchnk/auth_service/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CheckUser(login, password string) (models.User, error)
}

type User interface {
	GetUserById(id int64) (models.User, error)
	DeleteUserById(id int64) (string, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user models.User) (string, error)
}

type Service struct {
	Authorization
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User: NewUserService(repos.User),
	}
}
