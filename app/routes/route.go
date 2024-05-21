package routes

import (
	"project-2/app/middlewares"
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

	//userHandler
	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.PUT("/users/:id", userHandlerAPI.Update, middlewares.JWTMiddleware())

	//roomHandler
	e.POST("/rooms", userHandlerAPI.Register)
}
