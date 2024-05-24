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

func (m *model) GetReservationHistory(userID uint) ([]reservations.Reservation, error) {
    var reservations []reservations.Reservation
    if err := m.db.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
        return nil, err
    }
    return reservations, nil
}
