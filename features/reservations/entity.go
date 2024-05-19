package reservations

import "time"

type Reservation struct {
	ReservationID    string    `json:"reservation_id"`
	UserID           string    `json:"user_id"`
	RoomID           string    `json:"room_id"`
	BookingStartdate time.Time `json:"booking_start_date"`
	BookingEnddate   time.Time `json:"booking_end_date"`
	QuantityGuest    int       `json:"quantity_guest"`
	QuantityNights   int       `json:"quantity_nights"`
	AirbnbFee        int       `json:"airbnb_fee"`
	TotalPrice       int       `json:"total_price"`
	PaymentStatus    string    `json:"payment_status"`
	XenditInvoiceID  string    `json:"xendit_invoice"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
}

func (r Reservation) IsValidStatus() bool {
	return r.PaymentStatus == "Success" || r.PaymentStatus == "Failed"
}
