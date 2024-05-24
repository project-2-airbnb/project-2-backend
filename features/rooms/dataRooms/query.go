package datarooms

import (
	"log"
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
	// Begin a transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Convert rooms.Room to datarooms.Rooms
	newRoom := Rooms{
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

	// Create the room
	if err := tx.Create(&newRoom).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create the room facilities
	for _, facility := range room.Facilities {
		roomFacility := RoomFacilitys{
			RoomID:     newRoom.ID,
			FacilityID: facility.FacilityID,
		}
		if err := tx.Create(&roomFacility).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
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
	// Begin a transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()

	// Find the room to update
	var existingRoom Rooms
	if err := tx.First(&existingRoom, roomid).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update fields
	existingRoom.UserID = room.UserID
	existingRoom.RoomPicture = room.RoomPicture
	existingRoom.RoomName = room.RoomName
	existingRoom.Description = room.Description
	existingRoom.Location = room.Location
	existingRoom.QuantityGuest = room.QuantityGuest
	existingRoom.QuantityBedroom = room.QuantityBedroom
	existingRoom.QuantityBed = room.QuantityBed
	existingRoom.Price = room.Price

	// Save the changes
	if err := tx.Save(&existingRoom).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete existing room facilities
	if err := tx.Where("room_id = ?", roomid).Delete(&RoomFacilitys{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create the room facilities
	for _, facility := range room.Facilities {
		roomFacility := RoomFacilitys{
			RoomID:     roomid,
			FacilityID: facility.FacilityID,
		}
		if err := tx.Create(&roomFacility).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
}

// GetAllRooms implements rooms.DataRoominterface.
func (r *roomQuery) GetAllRooms() ([]rooms.Room, error) {
	var roomsList []Rooms
	result := r.db.Preload("RoomFacilitas.Facility").Find(&roomsList)
	if result.Error != nil {
		return nil, result.Error
	}

	// Fill FacilityNames for each room
	for i := range roomsList {
		roomsList[i].FacilityNames = extractFacilityNames(roomsList[i].RoomFacilitas)
	}

	// Convert Rooms to rooms.Room
	var resultRoomsList []rooms.Room
	for _, room := range roomsList {
		resultRoomsList = append(resultRoomsList, rooms.Room{
			UserID:          room.UserID,
			RoomPicture:     room.RoomPicture,
			RoomName:        room.RoomName,
			QuantityGuest:   room.QuantityGuest,
			QuantityBedroom: room.QuantityBedroom,
			QuantityBed:     room.QuantityBed,
			Price:           room.Price,
			FacilityNames:   room.FacilityNames,
		})
	}
	return resultRoomsList, nil
}

func extractFacilityNames(roomFacilities []RoomFacilitys) []string {
	var facilityNames []string
	for _, rf := range roomFacilities {
		facilityNames = append(facilityNames, rf.Facility.FacilityName)
	}
	return facilityNames
}

// GetRoomByName implements rooms.DataRoominterface.
func (r *roomQuery) GetRoomByName(roomName string) ([]rooms.Room, error) {
	var roomsList []Rooms
	tx := r.db.Preload("RoomFacilitas.Facility").Where("room_name = ?", roomName).First(&roomsList)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Fill FacilityNames for each room
	for i := range roomsList {
		roomsList[i].FacilityNames = extractFacilityNames(roomsList[i].RoomFacilitas)
	}

	// Convert Rooms to rooms.Room
	var resultRoomsList []rooms.Room
	for _, room := range roomsList {
		resultRoomsList = append(resultRoomsList, rooms.Room{
			UserID:          room.UserID,
			RoomPicture:     room.RoomPicture,
			RoomName:        room.RoomName,
			Description:     room.Description,
			Location:        room.Location,
			QuantityGuest:   room.QuantityGuest,
			QuantityBedroom: room.QuantityBedroom,
			QuantityBed:     room.QuantityBed,
			Price:           room.Price,
			FacilityNames:   room.FacilityNames,
		})
	}
	log.Println("roomsList: ", resultRoomsList)
	return resultRoomsList, nil

}

// GetRoomByID implements rooms.DataRoominterface.
func (r *roomQuery) GetRoomByID(roomID uint) (*rooms.Room, error) {
	var roomGorm Rooms
	tx := r.db.Preload("RoomFacilitas.Facility").First(&roomGorm, roomID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Mapping data dari roomGorm ke roomCore
	var roomCore rooms.Room
	roomCore.RoomID = roomGorm.ID
	roomCore.RoomPicture = roomGorm.RoomPicture
	roomCore.RoomName = roomGorm.RoomName
	roomCore.Description = roomGorm.Description
	roomCore.Location = roomGorm.Location
	roomCore.QuantityGuest = roomGorm.QuantityGuest
	roomCore.QuantityBedroom = roomGorm.QuantityBedroom
	roomCore.QuantityBed = roomGorm.QuantityBed
	roomCore.Price = roomGorm.Price

	return &roomCore, nil
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

// FacilitybyID implements rooms.DataRoominterface.
func (r *roomQuery) FacilitybyID(roomID uint) ([]rooms.Facility, error) {
	var facilityGorm []RoomFacilitys
	tx := r.db.Preload("Facility").Where("room_id = ?", roomID).Find(&facilityGorm)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var facilityCore []rooms.Facility

	for _, facility := range facilityGorm {
		facilityCore = append(facilityCore, rooms.Facility{
			FacilityID:   facility.FacilityID,
			FacilityName: facility.Facility.FacilityName,
		})
	}

	return facilityCore, nil
}
