package rooms

import "time"

type Room struct {
	RoomID          uint
	UserID          uint
	RoomPicture     string
	RoomName        string
	FullName        string
	Description     string
	Location        string
	QuantityGuest   int     `json:"quantity_guest"`
	QuantityBedroom int     `json:"quantity_bedroom"`
	QuantityBed     int     `json:"quantity_bed"`
	Price           int     `json:"price"`
	Rating          float32 `json:"rating"`
	Facilities      []Facility
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
	FacilityNames   []string  `gorm:"-" json:"facility_names"`
}

type Facility struct {
	FacilityID   uint   `json:"facility_id"`
	FacilityName string `json:"facility_name"`
}

type RoomFacility struct {
	FacilityID uint `json:"facility_id"`
	RoomID     uint `json:"room_id"`
}

type DataRoominterface interface {
	CreateRoom(room Room) error
	UpdateRoom(roomid uint, room Room) error
	DeleteRoom(roomid uint) error
	GetAllRooms() ([]Room, error)
	GetRoomByName(roomName string) ([]Room, error)
	GetRoomByID(roomID uint) (*Room, error)
	SelectByUserID(userID uint) (*Room, error)
	FacilitybyID(facilityID uint) ([]Facility, error)
}

type DataRoomService interface {
	AddRoom(room Room) error
	UpdateRoom(roomid uint, userid uint, room Room) error
	DeleteRoom(roomid uint, userid uint) error
	GetAllRooms(roomName string) ([]Room, error)
	GetRoomByID(roomID uint) (*Room, error)
}
