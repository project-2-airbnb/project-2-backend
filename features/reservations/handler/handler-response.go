package handler

type ReservationResponse struct {
	RoomID          uint   `json:"room_id"`
	CheckInDate     string `json:"check_in_date"`
	CheckOutDate    string `json:"check_out_date"`
	QuantityGuest   int    `json:"quantity_guest"`
	QuantityNights  int    `json:"quantity_night"`
	BiayaKebersihan int    `json:"biaya_kebersihan"`
	Pajak           int    `json:"pajak"`
	Total           int    `json:"total"`
	PaymentMethod   string `json:"payment_method"`
}