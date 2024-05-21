package datarooms

import (
	"project-2/features/rooms"

	"gorm.io/gorm"
)

type roomQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) rooms.DataRoominterface {
	return &roomQuery{
		db: db,
	}
}

// CreateRoom implements rooms.DataRoominterface.
func (r *roomQuery) CreateRoom(room rooms.Room) error {
	roomsGorm := Rooms{
		UserID:          room.UserID,
		RoomPicture:     room.RoomPicture,
		RoomName:        room.RoomName,
		Description:     room.Description,
		Location:        room.Location,
		QuantityGuest:   room.QuantityGuest,
		QuantityBedroom: room.QuantityBedroom,
		QuantityBed:     room.QuantityBed,
		Price:           room.Price,
	}
	tx := r.db.Create(&roomsGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteRoom implements rooms.DataRoominterface.
func (*roomQuery) DeleteRoom(roomid uint) error {
	panic("unimplemented")
}

// UpdateRoom implements rooms.DataRoominterface.
func (*roomQuery) UpdateRoom(room rooms.Room) (rooms.Room, error) {
	panic("unimplemented")
}

// GetUserByID implements rooms.DataRoominterface.
func (r *roomQuery) GetUserByID(userID uint) (*rooms.Room, error) {
	var roomGorm Rooms
	tx := r.db.First(&roomGorm, userID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping
	var roomcore = rooms.Room{
		UserID:          roomGorm.UserID,
		RoomPicture:     roomGorm.RoomPicture,
		RoomName:        roomGorm.RoomName,
		Description:     roomGorm.Description,
		Location:        roomGorm.Location,
		QuantityGuest:   roomGorm.QuantityGuest,
		QuantityBedroom: roomGorm.QuantityBedroom,
		QuantityBed:     roomGorm.QuantityBed,
		Price:           roomGorm.Price,
	}

	return &roomcore, nil
}
