package rooms

import "time"

type Room struct {
	RoomID          uint   `json:"room_id"`
	UserID          uint   `json:"user_id"`
	RoomPicture     string `json:"room_picture"`
	RoomName        string `json:"room_name"`
	Description     string `json:"description"`
	QuantityGuest   int    `json:"quantity_guest"`
	QuantityBedroom int    `json:"quantity_bedroom"`
	QuantityBath    int    `json:"quantity_bathroom"`
	Price           int    `json:"price"`
	RoomFacility    []RoomFacility
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type RoomFacility struct {
	FacilityID     uint `json:"facility_id"`
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
