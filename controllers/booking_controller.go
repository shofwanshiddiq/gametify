package controllers

import (
	"database/sql"
	"gametify/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookingController struct {
	DB *gorm.DB
}

func NewBookingController(db *gorm.DB) *BookingController {
	return &BookingController{DB: db}
}

func (b *BookingController) GetAllBookings(c *gin.Context) {
	var bookings []models.Booking
	if err := b.DB.Find(&bookings).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, bookings)
}

func (b *BookingController) GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking
	if err := b.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		c.JSON(404, gin.H{"error": "Booking not found"})
		return
	}
	c.JSON(200, booking)
}

func (b *BookingController) CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if the room is already booked in the given time range
	var existingBooking models.Booking
	err := b.DB.
		Where("room_id = ? AND start_time < ? AND end_time > ?", booking.RoomId, booking.EndTime, booking.StartTime).
		First(&existingBooking).Error

	if err == nil {
		// There's a conflict
		c.JSON(400, gin.H{"error": "Time already taken for this room"})
		return
	} else if err != gorm.ErrRecordNotFound {
		// An actual DB error occurred
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Save booking if no conflict
	if err := b.DB.Create(&booking).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, booking)
}

func (b *BookingController) UpdateBookingStatus(c *gin.Context) {
	id := c.Param("id")

	// Cek apakah booking id ada
	var booking models.Booking
	if err := b.DB.First(&booking, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Booking not found"})
		return
	}

	// Pastikan json terdiri dari book_status
	var req struct {
		BookStatus models.BookStatus `json:"book_status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update book_status
	if err := b.DB.Model(&booking).Update("book_status", req.BookStatus).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":     "Booking status updated successfully",
		"booking_id":  booking.ID,
		"book_status": req.BookStatus,
	})
}

func (b *BookingController) PostBookingRating(c *gin.Context) {
	id := c.Param("id")

	// Check if booking exists
	var booking models.Booking
	if err := b.DB.First(&booking, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Booking not found"})
		return
	}

	// Check if booking status is COMPLETED
	if booking.BookStatus != models.StatusCompleted {
		c.JSON(400, gin.H{"error": "Only completed bookings can be rated"})
		return
	}

	// Parse and validate rating input
	var req struct {
		Rating int `json:"rating"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if req.Rating < 1 || req.Rating > 5 {
		c.JSON(400, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	// Save rating
	if err := b.DB.Model(&booking).Update("rating", req.Rating).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":    "Rating submitted successfully",
		"booking_id": booking.ID,
		"rating":     req.Rating,
	})
}

func (b *BookingController) GetAverageRoomRating(c *gin.Context) {
	roomID := c.Param("room_id")
	var avgRating sql.NullFloat64

	err := b.DB.
		Table("bookings").
		Select("AVG(rating)").
		Where("room_id = ?", roomID).
		Scan(&avgRating).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// if no rating exists yet
	if !avgRating.Valid {
		c.JSON(http.StatusOK, gin.H{"room_id": roomID, "average_rating": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"room_id":        roomID,
		"average_rating": avgRating.Float64,
	})
}

func (b *BookingController) GetAverageRatingByPlace(c *gin.Context) {
	placeID := c.Param("place_id") // Correct parameter name
	var avgRating sql.NullFloat64

	err := b.DB.
		Table("bookings").
		Select("AVG(bookings.rating)").
		Joins("JOIN rooms ON bookings.room_id = rooms.id").
		Where("rooms.place_id = ?", placeID).
		Where("bookings.rating IS NOT NULL").
		Scan(&avgRating).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !avgRating.Valid {
		c.JSON(http.StatusOK, gin.H{
			"place_id":     placeID,
			"average_rate": nil,
			"message":      "No ratings available",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"place_id":     placeID,
		"average_rate": avgRating.Float64,
	})
}
