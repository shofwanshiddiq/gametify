package controllers

import (
	"fmt"
	"gametify/models"
	"gametify/services"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

// func (uc *UserController) GetAllUsers(c *gin.Context) {
// 	users, err := uc.service.GetAllUsers()
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, users)
// }

func (uc *UserController) GetAllUsers(c *gin.Context) {

	role, exists := c.Get("user_role")
	if !exists || role != models.TypeAdmin {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	users, err := uc.service.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {

	role, exists := c.Get("user_role")
	if !exists || role != models.TypeAdmin {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := uc.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {

	role, exists := c.Get("user_role")
	if !exists || role != models.TypeAdmin {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	if name, ok := updateData["name"].(string); ok && name != "" {
		user.Name = name
	}
	if email, ok := updateData["email"].(string); ok && email != "" {
		user.Email = email
	}

	updated, err := uc.service.UpdateUser(uint(id), user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updated)
}

func (uc *UserController) DeleteUser(c *gin.Context) {

	role, exists := c.Get("user_role")
	if !exists || role != models.TypeAdmin {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := uc.service.DeleteUser(uint(id)); err != nil {
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

	// Validate size (max 10MB)
	if file.Size > 10<<20 {
		c.JSON(400, gin.H{"error": "File is too large. Max 10MB allowed"})
		return
	}

	// Validate type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		c.JSON(400, gin.H{"error": "Invalid file type. Only PNG, JPG, JPEG allowed"})
		return
	}

	savePath := fmt.Sprintf("uploads/profile_%d_%s", userID, file.Filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	if err := uc.service.UploadProfilePicture(userID, savePath); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile picture uploaded successfully",
		"image":   "/" + savePath,
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

	picPath, err := uc.service.GetProfilePicture(userID)
	if err != nil || picPath == "" {
		c.JSON(404, gin.H{"error": "No profile picture found"})
		return
	}

	c.JSON(200, gin.H{"profile_picture": "/" + picPath})
}

func (uc *UserController) GetProfileData(c *gin.Context) {
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

	user, err := uc.service.GetUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
