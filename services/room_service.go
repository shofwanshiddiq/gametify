package services

import (
	"gametify/models"
	"gametify/repositories"
)

type RoomService interface {
	GetAllPlaces() ([]models.Place, error)
	GetPlaceByID(id string) (*models.Place, error)
	GetAllRooms() ([]models.Room, error)
	GetRoomByID(id string) (*models.Room, error)
	GetRoomsByPlaceID(placeID string) ([]models.Room, error)
	GetRoomsByConsoleType(consoleType string) ([]models.Room, error)
}

type roomService struct {
	repo repositories.RoomRepository
}

func NewRoomService(repo repositories.RoomRepository) RoomService {
	return &roomService{repo}
}

func (s *roomService) GetAllPlaces() ([]models.Place, error) {
	return s.repo.FindAllPlaces()
}

func (s *roomService) GetPlaceByID(id string) (*models.Place, error) {
	return s.repo.FindPlaceByID(id)
}

func (s *roomService) GetAllRooms() ([]models.Room, error) {
	return s.repo.FindAllRooms()
}

func (s *roomService) GetRoomByID(id string) (*models.Room, error) {
	return s.repo.FindRoomByID(id)
}

func (s *roomService) GetRoomsByPlaceID(placeID string) ([]models.Room, error) {
	return s.repo.FindRoomsByPlaceID(placeID)
}

func (s *roomService) GetRoomsByConsoleType(consoleType string) ([]models.Room, error) {
	return s.repo.FindRoomsByConsoleType(consoleType)
}
