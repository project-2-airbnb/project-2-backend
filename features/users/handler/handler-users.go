package handler

import (
	"project-2/features/users"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.ServiceUserInterface
}

func New(us users.ServiceUserInterface) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	panic("not implemented")
}
