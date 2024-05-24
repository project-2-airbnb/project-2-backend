package reservations

import "time"

type Reservation struct {
	ReservationID   uint    
	UserID          uint    
	RoomID          uint    
	CheckInDate     string
	CheckOutDate    string 
	QuantityGuest   int       
	QuantityNights  int       
	BiayaKebersihan int       
	Pajak           int       
	Total           int       
	PaymentMethod   string    
	CreatedAt       time.Time 
	UpdatedAt       time.Time 
	DeletedAt       time.Time 
}

type ReservationModel interface {
	AddReservation(reservation Reservation) error
}

type ReservationService interface {
	AddReservation(reservation Reservation) error
}
