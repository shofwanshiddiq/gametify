package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gametify/config"
	"gametify/models"
	"gametify/repositories"
	"gametify/route"
	"gametify/services"
)

/*
GAMETIFY FINAL PROJECT

-- Installation --
go get -u github.com/gin-gonic/gin -- Gin Gonic
go get -u gorm.io/gorm -- Gorm
go get -u gorm.io/driver/mysql -- Driver MySQL
go get github.com/golang-jwt/jwt/v5 -- JWT Authentication
go get github.com/joho/godotenv -- ENV
go get github.com/golang-jwt/jwt/v5 -- JWT
*/

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file")
	}

	db := config.ConnectDatabase()

	db.AutoMigrate(
		&models.User{},
		&models.Place{},
		&models.Room{},
		&models.Booking{},
		&models.Package{},
	)

	r := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	authRepo := repositories.NewAuthRepository(db)
	roomRepo := repositories.NewRoomRepository(db)
	bookingRepo := repositories.NewBookingRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(*authRepo)
	roomService := services.NewRoomService(roomRepo)
	bookingService := services.NewBookingService(bookingRepo)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gametify API!",
		})
	})

	route.SetupRoutes(r, userService, authService, roomService, bookingService, userRepo)

	// Start server
	port := ":8080"
	fmt.Printf("✅ Server is running on http://localhost%s\n", port)
	r.Run(port)
}
