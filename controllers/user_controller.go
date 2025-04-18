package controllers

import (
	"fmt"
	"path/filepath"
	"strings"

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

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := uc.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := uc.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Cek id user apakah ada di database
	if err := uc.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	// delete user berdasarkan id
	if err := uc.DB.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) UploadProfilePicture(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	file, err := c.FormFile("picture")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file is uploaded"})
		return
	}

	// validasi max 10 MB
	const maxSize = 10 << 20
	if file.Size > maxSize {
		c.JSON(400, gin.H{"error": "File is too large. Max 10MB allowed"})
		return
	}

	// validasi file type
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		c.JSON(400, gin.H{"error": "Invalid file type. Only PNG, JPG, JPEG allowed"})
		return
	}

	// Save profile picture ke folder
	savePath := fmt.Sprintf("uploads/profile_%d_%s", userID, file.Filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	// Update user profile picture di DB
	if err := uc.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("profile_picture", savePath).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile picture uploaded successfully",
		"image":   "/" + savePath, // if you serve static files
	})
}

func (uc *UserController) GetProfilePicture(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Check if profile_picture exists
	if user.ProfilePicture == "" {
		c.JSON(404, gin.H{"error": "No profile picture found"})
		return
	}

	c.JSON(200, gin.H{
		"profile_picture": "/" + user.ProfilePicture,
	})
}
