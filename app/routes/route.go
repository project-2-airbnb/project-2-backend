package routes

import (
	"project-2/app/middlewares"
	_reservationData "project-2/features/reservations/dataReservations"
	_reservationHandler "project-2/features/reservations/handler"
	_reservationService "project-2/features/reservations/service"
	_reviewData "project-2/features/review/dataReview"
	_reviewHandler "project-2/features/review/handler"
	_reviewService "project-2/features/review/service"
	_roomsData "project-2/features/rooms/dataRooms"
	_roomsHandler "project-2/features/rooms/handler"
	_roomsService "project-2/features/rooms/service"
	_userData "project-2/features/users/dataUsers"
	_userHandler "project-2/features/users/handler"
	_userService "project-2/features/users/service"
	"project-2/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)

	roomsData := _roomsData.New(db)
	roomsService := _roomsService.New(roomsData)
	roomsHandlerAPI := _roomsHandler.New(roomsService)

	Review := _reviewData.New(db)
	reviewService := _reviewService.New(Review)
	reviewHandlerAPI := _reviewHandler.New(reviewService)

	Reservation := _reservationData.New(db)
	reservationService := _reservationService.New(Reservation)
	reservationHandlerAPI := _reservationHandler.New(reservationService)
	



	//userHandler
	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.GET("/users", userHandlerAPI.GetProfile, middlewares.JWTMiddleware())

	//roomHandler
	e.POST("/rooms", roomsHandlerAPI.Create, middlewares.JWTMiddleware())
	e.DELETE("/rooms/:id", roomsHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.GET("/rooms", roomsHandlerAPI.AllRoom)
	e.GET("/rooms/:id", roomsHandlerAPI.GetRoomByID, middlewares.JWTMiddleware())
	e.PUT("/rooms/:id", roomsHandlerAPI.UpdateRoom, middlewares.JWTMiddleware())

	e.POST("reviews", reviewHandlerAPI.AddReview, middlewares.JWTMiddleware())

	e.POST("reservations", reservationHandlerAPI.AddReservation, middlewares.JWTMiddleware())
	e.GET("/reservations/history", reservationHandlerAPI.GetReservationHistory)
}
