package service

import (
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/amrchnk/auth_service/pkg/repository"
)

type Authorization interface {
	CheckUser(login, password string) (string, error)
	CreateUser(user models.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
