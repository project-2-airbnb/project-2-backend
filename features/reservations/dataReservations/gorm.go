package datareservations

import (
	"gorm.io/gorm"
)

type Reservations struct {
	gorm.Model
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
}
