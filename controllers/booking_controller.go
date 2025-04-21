package controllers

import (
	"gametify/models"
	"gametify/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService services.BookingService
}

func NewBookingController(bookingService services.BookingService) *BookingController {
	return &BookingController{bookingService}
}

func (b *BookingController) GetAllBookings(c *gin.Context) {
	bookings, err := b.bookingService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, bookings)
}

func (b *BookingController) GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	booking, err := b.bookingService.GetByID(id)
	if err != nil {
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

	if err := b.bookingService.CreateBooking(&booking); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, booking)
}

func (b *BookingController) UpdateBookingStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		BookStatus models.BookStatus `json:"book_status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := b.bookingService.UpdateStatus(id, req.BookStatus); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":     "Booking status updated successfully",
		"booking_id":  id,
		"book_status": req.BookStatus,
	})
}

func (b *BookingController) PostBookingRating(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Rating int `json:"rating"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := b.bookingService.RateBooking(id, req.Rating); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":    "Rating submitted successfully",
		"booking_id": id,
		"rating":     req.Rating,
	})
}

func (b *BookingController) GetAverageRoomRating(c *gin.Context) {
	roomID := c.Param("room_id")
	rating, err := b.bookingService.GetAverageRatingByRoom(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"room_id":        roomID,
		"average_rating": rating.Float64,
	})
}

func (b *BookingController) GetAverageRatingByPlace(c *gin.Context) {
	placeID := c.Param("place_id")
	rating, err := b.bookingService.GetAverageRatingByPlace(placeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"place_id":     placeID,
		"average_rate": rating.Float64,
	})
}
