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
	UserId     int        `json:"user_id"`
	RoomId     int        `json:"room_id"`
	StartTime  time.Time  `json:"start_time" gorm:"type:time"`
	EndTime    time.Time  `json:"end_time" gorm:"type:time"`
	TotalPrice float32    `json:"total_price"`
	BookStatus BookStatus `json:"book_status" gorm:"type:varchar(20)"`
	Rating     float32    `json:"rating"`
}
