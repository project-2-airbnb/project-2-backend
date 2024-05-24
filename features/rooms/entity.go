package rooms

import "time"

type Room struct {
	RoomID          uint       `json:"room_id"`
	UserID          uint       `json:"user_id"`
	RoomPicture     string     `json:"room_picture"`
	RoomName        string     `json:"room_name"`
	FullName        string     `json:"full_name"`
	Description     string     `json:"description"`
	Location        string     `json:"location"`
	QuantityGuest   int        `json:"quantity_guest"`
	QuantityBedroom int        `json:"quantity_bedroom"`
	QuantityBed     int        `json:"quantity_bed"`
	Price           int        `json:"price"`
	Rating          float32    `json:"rating"`
	Facilities      []Facility `gorm:"many2many:room_facilities;" json:"facilities"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       time.Time  `json:"deleted_at"`
	FacilityNames   []string   `gorm:"-" json:"facility_names"`
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
	GetRoomByName(roomName string) (*Room, error)
	GetRoomByID(roomID uint) (*Room, error)
	SelectByUserID(userID uint) (*Room, error)
}

type DataRoomService interface {
	AddRoom(room Room) error
	UpdateRoom(roomid uint, userid uint, room Room) error
	DeleteRoom(roomid uint, userid uint) error
	GetAllRooms(roomName string) ([]*Room, error)
	GetRoomByID(roomID uint) (*Room, error)
	GetUserRooms(userID uint) (*Room, error)
}
