package repositories

import (
	"database/sql"
	"gametify/models"
	"time"

	"gorm.io/gorm"
)

type BookingRepository interface {
	FindAll() ([]models.Booking, error)
	FindByID(id string) (*models.Booking, error)
	Create(booking *models.Booking) error
	UpdateStatus(id string, status models.BookStatus) error
	UpdateRating(id string, rating int) error
	IsTimeSlotTaken(roomID uint, start, end time.Time) (bool, error)
	GetAverageRoomRating(roomID string) (sql.NullFloat64, error)
	GetAverageRatingByPlace(placeID string) (sql.NullFloat64, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) FindAll() ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) FindByID(id string) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.First(&booking, id).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) Create(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) UpdateStatus(id string, status models.BookStatus) error {
	return r.db.Model(&models.Booking{}).Where("id = ?", id).Update("book_status", status).Error
}

func (r *bookingRepository) UpdateRating(id string, rating int) error {
	return r.db.Model(&models.Booking{}).Where("id = ?", id).Update("rating", rating).Error
}

func (r *bookingRepository) IsTimeSlotTaken(roomID uint, start, end time.Time) (bool, error) {
	var booking models.Booking
	err := r.db.Where("room_id = ? AND start_time < ? AND end_time > ?", roomID, end, start).
		First(&booking).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *bookingRepository) GetAverageRoomRating(roomID string) (sql.NullFloat64, error) {
	var avg sql.NullFloat64
	err := r.db.Table("bookings").Select("AVG(rating)").Where("room_id = ?", roomID).Scan(&avg).Error
	return avg, err
}

func (r *bookingRepository) GetAverageRatingByPlace(placeID string) (sql.NullFloat64, error) {
	var avg sql.NullFloat64
	err := r.db.Table("bookings").
		Select("AVG(bookings.rating)").
		Joins("JOIN rooms ON bookings.room_id = rooms.id").
		Where("rooms.place_id = ? AND bookings.rating IS NOT NULL", placeID).
		Scan(&avg).Error
	return avg, err
}
