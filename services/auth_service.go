package services

import (
	"gametify/models"
	"gametify/repositories"
	"gametify/utils"
)

type AuthService interface {
	Register(user *models.User) (string, error)
	Login(email, password string) (string, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(user *models.User) (string, error) {
	// Hash password
	if err := user.HashPassword(user.Password); err != nil {
		return "", err
	}

	// Save user to DB
	if err := s.repo.Create(user); err != nil {
		return "", err
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	// Check password
	if err := user.CheckPassword(password); err != nil {
		return "", err
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
