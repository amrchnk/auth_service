package service

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/amrchnk/auth_service/pkg/repository"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) CheckUser(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	userSession, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	return string(userSession), nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
