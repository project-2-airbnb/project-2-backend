package handler

import (
	"net/http"
	"project-2/features/users"
	"project-2/utils/responses"
	"strconv"
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

func (uh *UserHandler) Login(c echo.Context) error {
	// membaca data dari request body
	loginReq := LoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	// melakukan login
	_, token, err := uh.userService.LoginAccount(loginReq.Email, loginReq.Password, loginReq.UserType)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("login gagal: "+err.Error(), nil))
	}

	// mengembalikan respons dengan token
	return c.JSON(http.StatusOK, responses.JSONWebResponse("login berhasil", echo.Map{"token": token}))
}

func (uh *UserHandler) Update(c echo.Context) error {
	userid := c.Param("id")
	idConv, errConv := strconv.Atoi(userid)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error convert data: "+errConv.Error(), nil))
	}

	// membaca data dari request body
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	// Membaca file gambar pengguna (jika ada)
	file, err := c.FormFile("profile_picture")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Gagal membaca file gambar: "+err.Error(), nil))
	}

	// Jika file ada, unggah ke Cloudinary
	var imageURL string
	if file != nil {
		// Buka file
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Gagal membuka file gambar: "+err.Error(), nil))
		}
		defer src.Close()

		// Upload file ke Cloudinary
		imageURL, err = newUser.uploadToCloudinary(src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Gagal mengunggah gambar: "+err.Error(), nil))
		}
	}
	// mapping  dari request ke
	dataUser := users.User{
		FullName:       newUser.FullName,
		Email:          newUser.Email,
		Password:       newUser.Password,
		RetypePassword: newUser.RetypePassword,
		Address:        newUser.Address,
		PhoneNumber:    newUser.PhoneNumber,
		PictureProfile: imageURL,
	}
	// memanggil/mengirimkan data ke method service layer
	errInsert := uh.userService.UpdateProfile(uint(idConv), dataUser)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("perubahan data gagal: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("perubahan data gagal: "+errInsert.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("perubahan data berhasil", nil))
}
