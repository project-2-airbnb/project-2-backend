package datareservations

import (
	"project-2/features/reservations"

	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservations.ReservationModel {
	return &model{
		db:db,
	}
}

func (m *model) AddReservation(reservation reservations.Reservation) error {
	tx := m.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	newReservation := Reservations{
		UserID: reservation.UserID,
		RoomID: reservation.RoomID,
		CheckInDate: reservation.CheckInDate,
		CheckOutDate: reservation.CheckOutDate,
		QuantityGuest: reservation.QuantityGuest,
		BiayaKebersihan: reservation.BiayaKebersihan,
		Pajak: reservation.Pajak,
		Total: reservation.Total,
		PaymentMethod: reservation.PaymentMethod,
	}

	if err := tx.Create(&newReservation).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}