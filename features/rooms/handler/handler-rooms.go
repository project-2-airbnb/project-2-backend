package handler

import (
	"errors"
	"log"
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

	// Read data from the request body
	newRoom := RoomRequest{}
	errBind := c.Bind(&newRoom)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
	}

	// Read user image file (if any)
	file, err := c.FormFile("room_picture")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Failed to read image file: "+err.Error(), nil))
	}

	// If a file exists, upload to Cloudinary
	var imageURL string
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to open image file: "+err.Error(), nil))
		}
		defer src.Close()

		imageURL, err = newRoom.uploadToCloudinary(src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to upload image: "+err.Error(), nil))
		}
	}

	// Construct the room data
	room := rooms.Room{
		UserID:          uint(userID),
		RoomPicture:     imageURL,
		RoomName:        newRoom.RoomName,
		Description:     newRoom.Description,
		Location:        newRoom.Location,
		QuantityGuest:   newRoom.QuantityGuest,
		QuantityBedroom: newRoom.QuantityBedroom,
		QuantityBed:     newRoom.QuantityBed,
		Price:           newRoom.Price,
		Facilities:      newRoom.toFacilities(),
	}
	// Add the room using the service layer
	err = rh.roomService.AddRoom(room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("membuat room gagal: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("membuat room berhasil", nil))
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

func (rh *RoomHandler) AllRoom(c echo.Context) error {
	roomName := c.QueryParam("roomname")

	// Panggil GetAllRooms dari service layer
	rooms, err := rh.roomService.GetAllRooms(roomName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal mendapatkan semua room: "+err.Error(), nil))
	}

	// Konversi []*rooms.Room ke []RoomResponse
	var roomResponses []RoomResponse
	for _, room := range rooms {
		// Convert facilities to slice of facility names
		var facilityNames []string
		for _, facility := range room.Facilities {
			facilityNames = append(facilityNames, facility.FacilityName)
		}

		roomResponse := RoomResponse{
			RoomPicture:     room.RoomPicture,
			RoomName:        room.RoomName,
			QuantityGuest:   room.QuantityGuest,
			QuantityBedroom: room.QuantityBedroom,
			QuantityBed:     room.QuantityBed,
			Price:           room.Price,
			Rating:          room.Rating,
			Facilities:      facilityNames,
		}
		roomResponses = append(roomResponses, roomResponse)
		log.Println("data fasilitas", roomResponses)
	}

	// Kirim respons JSON yang berisi data ruangan
	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil mendapatkan semua room", roomResponses))
}

func (rh *RoomHandler) GetRoomByID(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error convert data: "+errConv.Error(), nil))
	}

	// Call the room service to get room by ID
	room, err := rh.roomService.GetRoomByID(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get room by ID: "+err.Error(), nil))
	}

	// Convert room data to RoomResponse
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
		Facilities:      room.FacilityNames,
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil mendapatkan data", roomResponse))
}

func (rh *RoomHandler) UpdateRoom(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error convert data: "+errConv.Error(), nil))
	}

	// Membaca data dari body permintaan
	updatedRoom := RoomRequest{}
	errBind := c.Bind(&updatedRoom)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
	}

	// Membaca file gambar ruangan (jika ada)
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
		imageURL, err = updatedRoom.uploadToCloudinary(src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Gagal mengunggah gambar: "+err.Error(), nil))
		}
	}

	// Mapping request ke struct Room
	dataRoom := rooms.Room{
		RoomID:          uint(idConv),
		UserID:          uint(userID),
		RoomPicture:     imageURL,
		RoomName:        updatedRoom.RoomName,
		Description:     updatedRoom.Description,
		Location:        updatedRoom.Location,
		QuantityGuest:   updatedRoom.QuantityGuest,
		QuantityBedroom: updatedRoom.QuantityBedroom,
		QuantityBed:     updatedRoom.QuantityBed,
		Price:           updatedRoom.Price,
		Facilities:      updatedRoom.toFacilities(),
	}

	// Memanggil service layer untuk memperbarui data ruangan
	if err := rh.roomService.UpdateRoom(uint(idConv), uint(userID), dataRoom); err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("gagal memperbarui room: "+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal memperbarui room: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil memperbarui room", nil))

}

func (rh *RoomHandler) GetRoomByUserID(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	// Call the service layer to get rooms by user ID
	userRooms, err := rh.roomService.GetUserRooms(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get rooms by user ID: "+err.Error(), nil))
	}

	// Convert room data to RoomResponse
	roomResponse := RoomResponse{
		RoomPicture:     userRooms.RoomPicture,
		RoomName:        userRooms.RoomName,
		Description:     userRooms.Description,
		Location:        userRooms.Location,
		QuantityGuest:   userRooms.QuantityGuest,
		QuantityBedroom: userRooms.QuantityBedroom,
		QuantityBed:     userRooms.QuantityBed,
		Price:           userRooms.Price,
		Rating:          userRooms.Rating,
		Facilities:      userRooms.FacilityNames,
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil mendapatkan data", roomResponse))
}
