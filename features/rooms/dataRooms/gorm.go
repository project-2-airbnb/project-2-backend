package datarooms

import (
	datareservations "project-2/features/reservations/dataReservations"
	datareview "project-2/features/review/dataReview"
	"project-2/features/rooms"

	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	UserID          uint                            `json:"user_id"`
	RoomPicture     string                          `json:"room_picture"`
	RoomName        string                          `json:"room_name"`
	Description     string                          `json:"description"`
	Location        string                          `json:"location"`
	QuantityGuest   int                             `json:"quantity_guest"`
	QuantityBedroom int                             `json:"quantity_bedroom"`
	QuantityBed     int                             `json:"quantity_bed"`
	Price           int                             `json:"price"`
	RoomFacilitas   []RoomFacilitys                 `gorm:"foreignKey:RoomID"`
	Reservations    []datareservations.Reservations `gorm:"foreignKey:RoomID"`
	Reviews         []datareview.Reviews            `gorm:"foreignKey:RoomID"`
	FacilityNames   []string                        `gorm:"-" json:"facility_names"` // Menyimpan nama fasilitas
}

type RoomFacilitys struct {
	gorm.Model
	RoomID     uint       `json:"room_id"`
	Facility   Facilities `gorm:"foreignkey:FacilityID"`
	FacilityID uint       `json:"facility_id"`
}

type Facilities struct {
	gorm.Model
	FacilityName string `json:"facility_name"`
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
