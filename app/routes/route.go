package routes

import (
	"project-2/app/middlewares"
	_roomsData "project-2/features/rooms/dataRooms"
	_roomsHandler "project-2/features/rooms/handler"
	_roomsService "project-2/features/rooms/service"
	_userData "project-2/features/users/dataUsers"
	_userHandler "project-2/features/users/handler"
	_userService "project-2/features/users/service"
	"project-2/utils/encrypts"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// Middleware untuk menangani CORS
	e.Use(middleware.CORS())

	// Middleware untuk menghapus trailing slash
	e.Pre(middleware.RemoveTrailingSlash())

	//userHandler
	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.PUT("/users/:id", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.GET("/users/:id", userHandlerAPI.GetProfile, middlewares.JWTMiddleware())

	//roomHandler
	e.POST("/rooms", roomsHandlerAPI.Create, middlewares.JWTMiddleware())
	e.DELETE("/rooms/:id", roomsHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.GET("/rooms", roomsHandlerAPI.AllRoom)
	e.GET("/rooms/:id", roomsHandlerAPI.GetRoomByID, middlewares.JWTMiddleware())
	e.PUT("/rooms/:id", roomsHandlerAPI.UpdateRoom, middlewares.JWTMiddleware())
}
