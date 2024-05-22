package service

import (
	"errors"
	// "log"
	"project-2/features/rooms"
)

type RoomService struct {
	roomData rooms.DataRoominterface
}

func New(rd rooms.DataRoominterface) rooms.DataRoomService {
	return &RoomService{
		roomData: rd,
	}

}

// AddRoom implements rooms.DataRoomService.
func (r *RoomService) AddRoom(room rooms.Room, facilities []rooms.Facility) error {
	if room.RoomName == "" || room.Location == "" || room.Description == "" || room.Price == 0 || room.QuantityBedroom == 0 || room.QuantityGuest == 0 {
		return errors.New("[validation] roomname/location/description/roomprice/quantitybedroom/quantityguest tidak boleh kosong")
	}

	// Create room
	err := r.roomData.CreateRoom(room, facilities)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom implements rooms.DataRoomService.
func (r *RoomService) DeleteRoom(roomid uint, userid uint) error {
	err := r.roomData.DeleteRoom(roomid)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRoom implements rooms.DataRoomService.
func (*RoomService) UpdateRoom(room rooms.Room) (rooms.Room, error) {
	panic("unimplemented")
}
