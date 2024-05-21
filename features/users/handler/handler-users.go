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
	// Membaca data dari body permintaan
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
	}

	// Mapping request ke struct User
	dataUser := users.User{
		FullName:       newUser.FullName,
		Email:          newUser.Email,
		Password:       newUser.Password,
		RetypePassword: newUser.RetypePassword,
		Address:        newUser.Address,
		PhoneNumber:    newUser.PhoneNumber,
		UserType:       newUser.UserType,
	}

	// Memanggil service layer untuk menyimpan data
	errInsert := uh.userService.RegistrasiAccount(dataUser)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Registrasi gagal: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Registrasi gagal: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Registrasi berhasil", nil))
}
