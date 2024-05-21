package handler

import (
	"net/http"
	"project-2/app/middlewares"
	"project-2/features/rooms"
	"project-2/utils/responses"
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

	// Mapping request ke struct User
	dataRoom := rooms.Room{
		UserID:          uint(userID),
		RoomPicture:     newRoom.RoomPicture,
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
