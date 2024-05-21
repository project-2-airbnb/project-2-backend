package reservations

import "time"

type Reservation struct {
	ReservationID   string    `json:"reservation_id"`
	UserID          string    `json:"user_id"`
	RoomID          string    `json:"room_id"`
	CheckInDate     time.Time `json:"check_in_date"`
	CheckOutDate    time.Time `json:"check_out_date"`
	QuantityGuest   int       `json:"quantity_guest"`
	QuantityNights  int       `json:"quantity_nights"`
	BiayaKebersihan int       `json:"biaya_kebersihan"`
	Pajak           int       `json:"pajak"`
	Total           int       `json:"total"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
