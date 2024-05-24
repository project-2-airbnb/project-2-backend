package handler

type RoomResponse struct {
	RoomPicture     string   `json:"room_picture,omitempty" form:"room_picture"`
	RoomName        string   `json:"room_name,omitempty" form:"room_name"`
	FullName        string   `json:"full_name"`
	Description     string   `json:"description,omitempty" form:"description"`
	Location        string   `json:"location,omitempty" form:"location"`
	QuantityGuest   int      `json:"quantity_guest,omitempty" form:"quantity_guest"`
	QuantityBedroom int      `json:"quantity_bedroom,omitempty" form:"quantity_bedroom"`
	QuantityBed     int      `json:"quantity_bed,omitempty" form:"quantity_bed"`
	Price           int      `json:"price,omitempty" form:"price"`
	Rating          float32  `json:"rating" form:"rating"`
	Facilities      []string `json:"facilities" form:"facilities"`
}
