package controllers

import (
	"gametify/models"
	"gametify/services"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomService services.RoomService
}

func NewRoomController(roomService services.RoomService) *RoomController {
	return &RoomController{roomService}
}

func (r *RoomController) GetAllPlaces(c *gin.Context) {
	places, err := r.roomService.GetAllPlaces()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, places)
}

func (r *RoomController) GetPlaceByID(c *gin.Context) {
	id := c.Param("id")
	place, err := r.roomService.GetPlaceByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Place not found"})
		return
	}
	c.JSON(200, place)
}

func (r *RoomController) GetAllRooms(c *gin.Context) {
	rooms, err := r.roomService.GetAllRooms()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, rooms)
}

func (r *RoomController) GetRoomByID(c *gin.Context) {
	id := c.Param("id")
	room, err := r.roomService.GetRoomByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Room not found"})
		return
	}
	c.JSON(200, room)
}

func (r *RoomController) GetRoomsByPlaceID(c *gin.Context) {
	placeID := c.Param("place_id")
	rooms, err := r.roomService.GetRoomsByPlaceID(placeID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, rooms)
}

func (r *RoomController) GetRoomsByConsoleType(c *gin.Context) {
	consoleType := c.Param("console_type")
	rooms, err := r.roomService.GetRoomsByConsoleType(consoleType)
	if err != nil {
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
