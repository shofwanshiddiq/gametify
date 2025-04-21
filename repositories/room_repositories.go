package repositories

import (
	"gametify/models"

	"gorm.io/gorm"
)

type RoomRepository interface {
	FindAllPlaces() ([]models.Place, error)
	FindPlaceByID(id string) (*models.Place, error)
	FindAllRooms() ([]models.Room, error)
	FindRoomByID(id string) (*models.Room, error)
	FindRoomsByPlaceID(placeID string) ([]models.Room, error)
	FindRoomsByConsoleType(consoleType string) ([]models.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db}
}

func (r *roomRepository) FindAllPlaces() ([]models.Place, error) {
	var places []models.Place
	err := r.db.Preload("Rooms").Find(&places).Error
	return places, err
}

func (r *roomRepository) FindPlaceByID(id string) (*models.Place, error) {
	var place models.Place
	err := r.db.Preload("Rooms").Where("id = ?", id).First(&place).Error
	if err != nil {
		return nil, err
	}
	return &place, nil
}

func (r *roomRepository) FindAllRooms() ([]models.Room, error) {
	var rooms []models.Room
	err := r.db.Find(&rooms).Error
	return rooms, err
}

func (r *roomRepository) FindRoomByID(id string) (*models.Room, error) {
	var room models.Room
	err := r.db.Where("id = ?", id).First(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) FindRoomsByPlaceID(placeID string) ([]models.Room, error) {
	var rooms []models.Room
	err := r.db.Where("place_id = ?", placeID).Find(&rooms).Error
	return rooms, err
}

func (r *roomRepository) FindRoomsByConsoleType(consoleType string) ([]models.Room, error) {
	var rooms []models.Room
	err := r.db.Where("console_type = ?", consoleType).Find(&rooms).Error
	return rooms, err
}
