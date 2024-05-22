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
func (r *roomQuery) DeleteRoom(roomid uint) error {
	tx := r.db.Where("id = ?", roomid).Delete(&Rooms{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateRoom implements rooms.DataRoominterface.
func (r *roomQuery) UpdateRoom(roomid uint, room rooms.Room) error {
	var roomGorm Rooms
	tx := r.db.First(&roomGorm, roomid)
	if tx.Error != nil {
		return tx.Error
	}

	roomGorm.RoomName = room.RoomName
	roomGorm.Description = room.Description
	roomGorm.Location = room.Location
	roomGorm.QuantityGuest = room.QuantityGuest
	roomGorm.QuantityBedroom = room.QuantityBedroom
	roomGorm.QuantityBed = room.QuantityBed
	roomGorm.Price = room.Price
	roomGorm.RoomPicture = room.RoomPicture

	tx = r.db.Save(&roomGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// GetAllRooms implements rooms.DataRoominterface.
func (r *roomQuery) GetAllRooms() ([]rooms.Room, error) {
	var roomsList []rooms.Room
	result := r.db.Model(&rooms.Room{}).
		Select("rooms.*, COALESCE(AVG(reviews.rating), 0) AS rating").
		Joins("LEFT JOIN reviews ON reviews.room_id = rooms.id").
		Group("rooms.id").
		Find(&roomsList)
	if result.Error != nil {
		return nil, result.Error
	}
	return roomsList, nil
}

// GetRoomByName implements rooms.DataRoominterface.
func (r *roomQuery) GetRoomByName(roomName string) (*rooms.Room, error) {
	var roomsGorm Rooms
	if err := r.db.Where("room_name = ?", roomName).Find(&roomsGorm).Error; err != nil {
		return nil, err
	}
	// mapping
	var roomcore = rooms.Room{
		RoomPicture:     roomsGorm.RoomPicture,
		RoomName:        roomsGorm.RoomName,
		QuantityGuest:   roomsGorm.QuantityGuest,
		QuantityBedroom: roomsGorm.QuantityBedroom,
		QuantityBed:     roomsGorm.QuantityBed,
		Price:           roomsGorm.Price,
	}
	return &roomcore, nil
}

// GetRoomByID implements rooms.DataRoominterface.
func (r *roomQuery) GetRoomByID(roomID uint) (*rooms.Room, error) {
	var roomGorm Rooms
	tx := r.db.First(&roomGorm, roomID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping
	var roomcore = rooms.Room{
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

// SelectByUserID implements rooms.DataRoominterface.
func (r *roomQuery) SelectByUserID(userID uint) (*rooms.Room, error) {
	var userGorm Rooms
	tx := r.db.First(&userGorm, userID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping
	var usercore = rooms.Room{
		RoomID:          userGorm.ID,
		UserID:          userGorm.UserID,
		RoomPicture:     userGorm.RoomPicture,
		RoomName:        userGorm.RoomName,
		Description:     userGorm.Description,
		Location:        userGorm.Location,
		QuantityGuest:   userGorm.QuantityGuest,
		QuantityBedroom: userGorm.QuantityBedroom,
		QuantityBed:     userGorm.QuantityBed,
		Price:           userGorm.Price,
	}

	return &usercore, nil
}
