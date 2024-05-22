package handler

import (
	"errors"
	"net/http"
	"project-2/app/middlewares"
	"project-2/features/rooms"
	"project-2/utils/responses"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	roomService rooms.DataRoomService
}

func New(rh rooms.DataRoomService) *RoomHandler {
	return &RoomHandler{
		roomService: rh,
	}
}

func (rh *RoomHandler) Create(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	// Membaca data dari body permintaan
	newRoom := RoomRequest{}
	errBind := c.Bind(&newRoom)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
	}

	// Membaca file gambar pengguna (jika ada)
	file, err := c.FormFile("room_picture")
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
		imageURL, err = newRoom.uploadToCloudinary(src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Gagal mengunggah gambar: "+err.Error(), nil))
		}
	}

	// Mapping request ke struct User
	dataRoom := rooms.Room{
		UserID:          uint(userID),
		RoomPicture:     imageURL,
		RoomName:        newRoom.RoomName,
		Description:     newRoom.Description,
		Location:        newRoom.Location,
		QuantityGuest:   newRoom.QuantityGuest,
		QuantityBedroom: newRoom.QuantityBedroom,
		QuantityBed:     newRoom.QuantityBed,
		Price:           newRoom.Price,
	}

	// Memanggil service layer untuk menyimpan data
	if err := rh.roomService.AddRoom(dataRoom); err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("gagal membuat room: "+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal membuat room: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("berhasil membuat room", nil))
}

func (rh *RoomHandler) Delete(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error convert data: "+errConv.Error(), nil))
	}
	// Memanggil service layer untuk menghapus data
	if err := rh.roomService.DeleteRoom(uint(idConv), uint(userID)); err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("gagal menghapus room: "+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal menghapus room: "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil menghapus room", nil))
}

func (rh *RoomHandler) SearchRoomByname(c echo.Context) error {
	roomName := c.QueryParam("roomname")
	if roomName == "" {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Room name parameter is required", nil))
	}

	// Call the room service to search for rooms by name
	rooms, err := rh.roomService.GetRoomByName(roomName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("data room tidak ditemukan: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("pencarian berhasil", rooms))
}

func (rh *RoomHandler) AllRoom(c echo.Context) error {
	rooms, err := rh.roomService.GetAllRooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal mendapatkan semua room: "+err.Error(), nil))
	}

	// Konversi data Room menjadi RoomResponse
	roomResponses := make([]RoomResponse, 0)
	for _, room := range rooms {
		roomResponse := RoomResponse{
			RoomPicture:     room.RoomPicture,
			RoomName:        room.RoomName,
			Description:     room.Description,
			Location:        room.Location,
			QuantityGuest:   room.QuantityGuest,
			QuantityBedroom: room.QuantityBedroom,
			QuantityBed:     room.QuantityBed,
			Price:           room.Price,
			Rating:          room.Rating,
		}
		roomResponses = append(roomResponses, roomResponse)
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil mendapatkan semua room", roomResponses))
}
