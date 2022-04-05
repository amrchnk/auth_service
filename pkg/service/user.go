package service

import (
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/amrchnk/auth_service/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserById(id int64) (models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) DeleteUserById(id int64) (string, error) {
	return s.repo.DeleteUserById(id)
}

func (s *UserService) UpdateUser(user models.User) (string, error){
	return s.repo.UpdateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
