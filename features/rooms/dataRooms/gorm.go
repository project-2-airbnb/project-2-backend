package datarooms

import (
	datareservations "project-2/features/reservations/dataReservations"
	datareview "project-2/features/review/dataReview"
	"project-2/features/rooms"

	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	UserID          uint
	RoomPicture     string
	RoomName        string
	Description     string `gorm:"type:text"`
	Location        string
	QuantityGuest   int
	QuantityBedroom int
	QuantityBed     int
	Price           int
	RoomFacilitas   []RoomFacilitys                 `gorm:"foreignKey:RoomID"`
	Reservations    []datareservations.Reservations `gorm:"foreignKey:RoomID"`
	Reviews         []datareview.Reviews            `gorm:"foreignKey:RoomID"`
	Facilities      []Facilities                    `gorm:"many2many:room_facilities"`
}

type RoomFacilitys struct {
	gorm.Model
	RoomID     uint
	Facility   Facilities `gorm:"foreignkey:FacilityID"`
	FacilityID uint
}

type Facilities struct {
	gorm.Model
	FacilityName string `gorm:"type:text"`
}

func (r Rooms) ModelToRoom() rooms.Room {
	return rooms.Room{
		RoomID:          r.ID,
		UserID:          r.UserID,
		RoomPicture:     r.RoomPicture,
		RoomName:        r.RoomName,
		Description:     r.Description,
		Location:        r.Location,
		QuantityGuest:   r.QuantityGuest,
		QuantityBedroom: r.QuantityBedroom,
		QuantityBed:     r.QuantityBed,
		Price:           r.Price,
	}
}
