package datareservations

import (
	"time"

	"gorm.io/gorm"
)

type Reservations struct {
	gorm.Model
	UserID           uint      `json:"user_id"`
	RoomID           uint      `json:"room_id"`
	BookingStartdate time.Time `json:"booking_start_date"`
	BookingEnddate   time.Time `json:"booking_end_date"`
	QuantityGuest    int       `json:"quantity_guest"`
	QuantityNights   int       `json:"quantity_nights"`
	AirbnbFee        int       `json:"airbnb_fee"`
	TotalPrice       int       `json:"total_price"`
	PaymentStatus    string    `json:"payment_status"`
	XenditInvoiceID  string    `json:"xendit_invoice_id"`
}

func (r Reservations) IsValidStatus() bool {
	return r.PaymentStatus == "Success" || r.PaymentStatus == "Failed"
}
