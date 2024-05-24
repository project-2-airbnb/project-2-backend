package service

import (
	"errors"
	"project-2/features/reservations"
)

type ReservationService struct {
	reservationData reservations.ReservationModel
}

func New(rm reservations.ReservationModel) reservations.ReservationService {
	return &ReservationService{
		reservationData: rm,
	}
}

func (r *ReservationService) AddReservation(reservation reservations.Reservation) error {
	if reservation.RoomID == 0 || reservation.BiayaKebersihan == 0 || reservation.CheckInDate == "" || reservation.CheckOutDate == "" || reservation.QuantityGuest == 0 || reservation.QuantityNights == 0 || reservation.Pajak == 0 || reservation.Total == 0 || reservation.PaymentMethod == "" {
		return errors.New("[validation] roomid/checkindate/checkoutdate/quantityguest/quantitynights/biayakebersihan/pajak/total/paymentmethod tidak boleh kosong")
	}

	err := r.reservationData.AddReservation(reservation)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReservationService) GetReservationsHistory(userID uint) ([]reservations.Reservation, error) {
    bookings, err := r.reservationData.GetReservationHistory(userID)
    if err != nil {
        return nil, err
    }
    return bookings, nil
}
