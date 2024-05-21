package datareservations

import (
	"time"

	"gorm.io/gorm"
)

type Reservations struct {
	gorm.Model
	UserID          uint      `json:"user_id"`
	RoomID          uint      `json:"room_id"`
	CheckInDate     time.Time `json:"check_in_date"`
	CheckOutDate    time.Time `json:"check_out_date"`
	QuantityGuest   int       `json:"quantity_guest"`
	QuantityNights  int       `json:"quantity_nights"`
	BiayaKebersihan int       `json:"biaya_kebersihan"`
	Pajak           int       `json:"pajak"`
	Total           int       `json:"total"`
	PaymentMethod   string    `json:"payment_method"`
}
