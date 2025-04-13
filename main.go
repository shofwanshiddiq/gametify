package main

import (
	"fmt"
	"gametify/config"
	"gametify/controllers"
	"gametify/middleware"
	"gametify/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	// Load .env file configuration
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}

	// Initialize Gin router
	r := gin.Default()

	// Connect to the database
	db := config.ConnectDatabase()

	// Auto migrate the models to database
	db.AutoMigrate(&models.User{}, &models.Place{}, &models.Rating{}, &models.Room{}, &models.Booking{}, &models.Package{})

	//  route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gametify API!",
			"routes": []string{
				"POST /login",
				"POST /register",
				"GET /rooms",
				"POST /book",
			},
		})
	})

	authController := controllers.NewAuthController(db)

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
		protected.GET("/users", controllers.NewUserController(db).GetAllUsers)
		protected.GET("/users/:id", controllers.NewUserController(db).GetUserByID)
	}

	// Start server	on port 8080
	port := ":8080"
	fmt.Println("✅ Server is running on http://localhost" + port)
	r.Run(port)
}
