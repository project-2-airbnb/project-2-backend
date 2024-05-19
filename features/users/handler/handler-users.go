package handler

import (
	"net/http"
	"project-2/features/users"
	"project-2/utils/responses"
	"strings"

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
	// membaca data dari request body
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	// mapping  dari request ke
	dataUser := users.User{
		UserPicture: newUser.UserPicture,
		UserName:    newUser.UserName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		Phone:       newUser.Phone,
		Address:     newUser.Address,
		UserType:    newUser.UserType,
	}
	// memanggil/mengirimkan data ke method service layer
	errInsert := uh.userService.RegistrasiAccount(dataUser)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("failed registration: "+errInsert.Error(), nil))

		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("registrasi gagal: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("registrasi berhasil", nil))

}
