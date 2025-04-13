package models

import (
	"time"

	"gorm.io/gorm"
)

/*
book_id
user_id
room_id
start_time
end_time
total_price
book_status
*/

// BookStatus is a custom type for booking status
type BookStatus string

// Enum values for BookStatus
const (
	StatusPending   BookStatus = "SUBMITTED"
	StatusConfirmed BookStatus = "CONFIRMED"
	StatusCanceled  BookStatus = "CANCELED"
	StatusCompleted BookStatus = "COMPLETED"
)

type Booking struct {
	gorm.Model
	BookingId  uint       `json:"booking_id"`
	UserId     int        `json:"user_id"`
	RoomId     int        `json:"room_id"`
	StartTime  time.Time  `json:"start_time"`
	EndTime    time.Time  `json:"end_time"`
	TotalPrice float32    `json:"total_price"`
	BookStatus BookStatus `json:"book_status" gorm:"type:varchar(20)"`
}

type Rating struct {
	gorm.Model
	RatingId  int     `json:"rating_id"`
	BookingId int     `json:"booking_id"`
	UserId    int     `json:"user_id"`
	PlaceId   int     `json:"place_id"`
	RoomId    int     `json:"room_id"`
	Rating    float32 `json:"rating"`
}
