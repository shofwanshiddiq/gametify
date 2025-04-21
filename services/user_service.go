package services

import (
	"gametify/models"
	"gametify/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, update models.User) (*models.User, error)
	DeleteUser(id uint) error
	UploadProfilePicture(userID uint, path string) error
	GetProfilePicture(userID uint) (string, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, update models.User) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Name = update.Name
	user.Email = update.Email
	err = s.repo.Update(user)
	return user, err
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) UploadProfilePicture(userID uint, path string) error {
	return s.repo.UpdateProfilePicture(userID, path)
}

func (s *userService) GetProfilePicture(userID uint) (string, error) {
	user, err := s.repo.FindByIDRaw(userID)
	if err != nil {
		return "", err
	}
	return user.ProfilePicture, nil
}
