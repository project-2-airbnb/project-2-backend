package datarooms

import (
	datareservations "project-2/features/reservations/dataReservations"
	datareview "project-2/features/review/dataReview"

	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	UserID          uint                            `json:"user_id"`
	RoomPicture     string                          `json:"room_picture"`
	RoomName        string                          `json:"room_name"`
	Description     string                          `json:"description"`
	QuantityGuest   int                             `json:"quantity_guest"`
	QuantityBedroom int                             `json:"quantity_bedroom"`
	QuantityBath    int                             `json:"quantity_bathroom"`
	Price           int                             `json:"price"`
	RoomFacility    []RoomFacilitys                 `gorm:"foreignKey:RoomID"`
	Reservations    []datareservations.Reservations `gorm:"foreignKey:RoomID"`
	Reviews         []datareview.Reviews            `gorm:"foreignKey:RoomID"`
}

type RoomFacilitys struct {
	gorm.Model
	RoomID         uint `json:"room_id"`
	Kitchen        bool `json:"kitchen"`
	Bathtub        bool `json:"bathtub"`
	Refrigerator   bool `json:"refrigerator"`
	Wifi           bool `json:"wifi"`
	TV             bool `json:"tv"`
	AirConditioner bool `json:"air_conditioner"`
	FreeParking    bool `json:"free_parking"`
	BeachView      bool `json:"beach_view"`
	ParkView       bool `json:"park_view"`
}
