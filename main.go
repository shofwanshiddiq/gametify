package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gametify/config"
	"gametify/controllers"
	"gametify/middleware"
	"gametify/models"
	"gametify/repositories"
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

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	roomController := controllers.NewRoomController(roomService)
	bookingController := controllers.NewBookingController(bookingService)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gametify API!",
		})
	})

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
	}

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", userController.GetAllUsers)
		protected.GET("/users/:id", userController.GetUserByID)
		protected.PUT("/users/:id", userController.UpdateUser)
		protected.DELETE("/users/:id", userController.DeleteUser)
		protected.POST("/users/profile-picture", userController.UploadProfilePicture)
		protected.GET("/users/profile-picture", userController.GetProfilePicture)

		protected.GET("/places", roomController.GetAllPlaces)
		protected.GET("/places/:id", roomController.GetPlaceByID)
		protected.GET("/rooms", roomController.GetAllRooms)
		protected.GET("/rooms/:id", roomController.GetRoomByID)
		protected.GET("/rooms/place/:place_id", roomController.GetRoomsByPlaceID)
		protected.GET("/rooms/console/:console_type", roomController.GetRoomsByConsoleType)
		protected.GET("/rooms/console", roomController.GetConsoleTypes)

		protected.GET("/bookings", bookingController.GetAllBookings)
		protected.GET("/bookings/:id", bookingController.GetBookingByID)
		protected.POST("/bookings", bookingController.CreateBooking)
		protected.PATCH("/bookings/:id", bookingController.UpdateBookingStatus)
		protected.POST("/bookings/:id/rate", bookingController.PostBookingRating)
		protected.GET("/bookings/room/:room_id", bookingController.GetAverageRoomRating)
		protected.GET("/bookings/place/:place_id", bookingController.GetAverageRatingByPlace)
	}

	// Start server
	port := ":8080"
	fmt.Printf("✅ Server is running on http://localhost%s\n", port)
	r.Run(port)
}
