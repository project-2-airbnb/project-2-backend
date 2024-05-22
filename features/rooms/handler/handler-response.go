package handler

type RoomResponse struct {
	RoomPicture     string  `json:"room_picture" form:"room_picture"`
	RoomName        string  `json:"room_name" form:"room_name"`
	Description     string  `json:"description" form:"description"`
	Location        string  `json:"location" form:"location"`
	QuantityGuest   int     `json:"quantity_guest" form:"quantity_guest"`
	QuantityBedroom int     `json:"quantity_bedroom" form:"quantity_bedroom"`
	QuantityBed     int     `json:"quantity_bed" form:"quantity_bed"`
	Price           int     `json:"price" form:"price"`
	Rating          float32 `json:"rating" form:"rating"`
}
