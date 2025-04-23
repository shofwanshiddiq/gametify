package route

import (
	"gametify/controllers"
	"gametify/middleware"
	"gametify/repositories"
	"gametify/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userService services.UserService, authService services.AuthService, roomService services.RoomService, bookingService services.BookingService, userRepo repositories.UserRepository) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			authController := controllers.NewAuthController(authService)
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
	}

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(userRepo)) // Pass userRepo to middleware
	{
		userController := controllers.NewUserController(userService)
		roomController := controllers.NewRoomController(roomService)
		bookingController := controllers.NewBookingController(bookingService)

		protected.GET("/users", userController.GetAllUsers)
		protected.GET("/users/:id", userController.GetUserByID)
		protected.GET("/users/profile", userController.GetProfileData)
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
}
