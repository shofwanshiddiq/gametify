package controllers

import (
	"gametify/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoomController struct {
	DB *gorm.DB
}

func NewRoomController(db *gorm.DB) *RoomController {
	return &RoomController{DB: db}
}

func (r *RoomController) GetAllPlaces(c *gin.Context) {
	var places []models.Place
	if err := r.DB.Find(&places).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, places)
}

func (r *RoomController) GetPlaceByID(c *gin.Context) {
	id := c.Param("id")
	var place models.Place
	if err := r.DB.Where("id = ?", id).First(&place).Error; err != nil {
		c.JSON(404, gin.H{"error": "Place not found"})
		return
	}
	c.JSON(200, place)
}

func (r *RoomController) GetAllRooms(c *gin.Context) {
	var rooms []models.Room
	if err := r.DB.Find(&rooms).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, rooms)
}

func (r *RoomController) GetRoomByID(c *gin.Context) {
	id := c.Param("id")
	var room models.Room
	if err := r.DB.Where("id = ?", id).First(&room).Error; err != nil {
		c.JSON(404, gin.H{"error": "Room not found"})
		return
	}
	c.JSON(200, room)
}

func (r *RoomController) GetRoomsByPlaceID(c *gin.Context) {
	placeID := c.Param("place_id")
	var rooms []models.Room
	if err := r.DB.Where("place_id = ?", placeID).Find(&rooms).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, rooms)
}

func (r *RoomController) GetRoomsByConsoleType(c *gin.Context) {
	consoleType := c.Param("console_type")
	var rooms []models.Room
	if err := r.DB.Where("console_type = ?", consoleType).Find(&rooms).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, rooms)
}

func (r *RoomController) GetConsoleTypes(c *gin.Context) {
	consoleTypes := []models.ConsoleType{
		models.TypePC,
		models.TypePS4,
		models.TypePS5,
	}
	c.JSON(200, consoleTypes)
}
