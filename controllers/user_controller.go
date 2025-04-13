package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gametify/models"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := uc.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}
