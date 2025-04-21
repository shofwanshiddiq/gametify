package services

import (
	"database/sql"
	"errors"
	"gametify/models"
	"gametify/repositories"
)

type BookingService interface {
	GetAll() ([]models.Booking, error)
	GetByID(id string) (*models.Booking, error)
	CreateBooking(booking *models.Booking) error
	UpdateStatus(id string, status models.BookStatus) error
	RateBooking(id string, rating int) error
	GetAverageRatingByRoom(roomID string) (sql.NullFloat64, error)
	GetAverageRatingByPlace(placeID string) (sql.NullFloat64, error)
}

type bookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(repo repositories.BookingRepository) BookingService {
	return &bookingService{repo}
}

func (s *bookingService) GetAll() ([]models.Booking, error) {
	return s.repo.FindAll()
}

func (s *bookingService) GetByID(id string) (*models.Booking, error) {
	return s.repo.FindByID(id)
}

func (s *bookingService) CreateBooking(booking *models.Booking) error {
	conflict, err := s.repo.IsTimeSlotTaken(uint(booking.RoomId), booking.StartTime, booking.EndTime)
	if err != nil {
		return err
	}
	if conflict {
		return errors.New("time already taken for this room")
	}
	return s.repo.Create(booking)
}

func (s *bookingService) UpdateStatus(id string, status models.BookStatus) error {
	return s.repo.UpdateStatus(id, status)
}

func (s *bookingService) RateBooking(id string, rating int) error {
	booking, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if booking.BookStatus != models.StatusCompleted {
		return errors.New("only completed bookings can be rated")
	}
	return s.repo.UpdateRating(id, rating)
}

func (s *bookingService) GetAverageRatingByRoom(roomID string) (sql.NullFloat64, error) {
	return s.repo.GetAverageRoomRating(roomID)
}

func (s *bookingService) GetAverageRatingByPlace(placeID string) (sql.NullFloat64, error) {
	return s.repo.GetAverageRatingByPlace(placeID)
}
