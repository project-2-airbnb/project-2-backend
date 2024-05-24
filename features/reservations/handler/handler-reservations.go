package handler

import (
	"net/http"
	"project-2/app/middlewares"
	"project-2/features/reservations"
	"project-2/utils/responses"

	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	ReservationService reservations.ReservationService
}

func New(rh reservations.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		ReservationService: rh,
	}
}

func (rh *ReservationHandler) AddReservation(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	newReservation := ReservationRequest{}
	errBind := c.Bind(&newReservation)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
	}

	reservation := reservations.Reservation{
		UserID: uint(userID),
		RoomID: newReservation.RoomID,
		CheckInDate: newReservation.CheckInDate,
		CheckOutDate: newReservation.CheckOutDate,
		QuantityGuest: newReservation.QuantityGuest,
		QuantityNights: newReservation.QuantityNights,
		BiayaKebersihan: newReservation.BiayaKebersihan,
		Pajak: newReservation.Pajak,
		Total: newReservation.Total,
		PaymentMethod: newReservation.PaymentMethod,
	}

	err := rh.ReservationService.AddReservation(reservation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal membuat reservasi: "+err.Error(),nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("membuat reservasi berhasil", nil))
}