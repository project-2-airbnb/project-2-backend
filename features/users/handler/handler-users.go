package handler

import (
	"net/http"
	"project-2/app/middlewares"
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
	_, token, err := uh.userService.LoginAccount(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("login gagal: "+err.Error(), nil))
	}

	// mengembalikan respons dengan token
	return c.JSON(http.StatusOK, responses.JSONWebResponse("login berhasil", echo.Map{"token": token}))
}

func (uh *UserHandler) Update(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
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
	errInsert := uh.userService.UpdateProfile(uint(userID), dataUser)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("perubahan data gagal: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("perubahan data gagal: "+errInsert.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("perubahan data berhasil", nil))
}

func (uh *UserHandler) Delete(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	errDelete := uh.userService.DeleteAccount(uint(userID))
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("hapus akun gagal: "+errDelete.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("hapus akun berhasil", nil))
}

func (uh *UserHandler) GetProfile(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	profile, err := uh.userService.GetProfile(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal mengambil profil: "+err.Error(), nil))
	}

	// Mapping data profil ke dalam format UserResponse
	userResponse := UserResponse{
		FullName:       profile.FullName,
		Email:          profile.Email,
		Address:        profile.Address,
		PhoneNumber:    profile.PhoneNumber,
		PictureProfile: profile.PictureProfile,
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil mengambil profil", userResponse))
}
