package rooms

import "time"

type Room struct {
	RoomID          uint      `json:"room_id"`
	UserID          uint      `json:"user_id"`
	RoomPicture     string    `json:"room_picture"`
	RoomName        string    `json:"room_name"`
	Description     string    `json:"description"`
	Location        string    `json:"location"`
	QuantityGuest   int       `json:"quantity_guest"`
	QuantityBedroom int       `json:"quantity_bedroom"`
	QuantityBed     int       `json:"quantity_bathroom"`
	Price           int       `json:"price"`
	Rating          float32   `json:"rating"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type RoomFacility struct {
	FacilityID uint `json:"facility_id"`
	RoomID     uint `json:"room_id"`
}

type Facility struct {
	FacilityID   uint   `json:"facility_id"`
	FacilityName string `json:"facility_name"`
}

type DataRoominterface interface {
	CreateRoom(room Room) error
	UpdateRoom(room Room) (Room, error)
	DeleteRoom(roomid uint) error
	GetAllRooms() ([]Room, error)
	GetRoomByName(roomName string) ([]Room, error)
}

type DataRoomService interface {
	AddRoom(room Room) error
	UpdateRoom(room Room) (Room, error)
	DeleteRoom(roomid uint, userid uint) error
	GetAllRooms() ([]Room, error)
	GetRoomByName(roomName string) ([]Room, error)
}
