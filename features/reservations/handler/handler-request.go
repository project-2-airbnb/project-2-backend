package handler

type ReservationRequest struct {
	RoomID          uint   `json:"room_id" form:"room_id"`
	CheckInDate     string `json:"check_in_date" form:"check_in_date"`
	CheckOutDate    string `json:"check_out_date" form:"check_out_date"`
	QuantityGuest   int    `json:"quantity_guest" form:"quantity_guest"`
	QuantityNights  int    `json:"quantity_night" form:"quantity_night"`
	BiayaKebersihan int    `json:"biaya_kebersihan" form:"biaya_kebersihan"`
	Pajak           int    `json:"pajak" form:"pajak"`
	Total           int    `json:"total" form:"total"`
	PaymentMethod   string `json:"payment_method" form:"payment_method"`
}