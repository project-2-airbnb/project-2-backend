package service

import (
	"errors"
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
func (r *RoomService) AddRoom(room rooms.Room) error {
	if room.RoomName == "" || room.Location == "" || room.Description == "" || room.Price == 0 || room.QuantityBedroom == 0 || room.QuantityGuest == 0 {
		return errors.New("[validation] roomname/location/description/roomprice/quantitybedroom/quantityguest tidak boleh kosong")
	}

	// Check if the user exists
	_, err := r.roomData.SelectByUserID(room.UserID)
	if err != nil {
		return errors.New("UserID tidak valid atau telah dihapus")
	}

	// Create room
	err = r.roomData.CreateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom implements rooms.DataRoomService.
func (r *RoomService) DeleteRoom(roomid uint, userid uint) error {
	if roomid <= 0 {
		return errors.New("invalid room ID")
	}
	cekuserid, err := r.roomData.SelectByUserID(roomid)
	if err != nil {
		return err
	}

	if cekuserid.UserID != userid {
		return errors.New("user id tidak sesuai")
	}

	return r.roomData.DeleteRoom(roomid)
}

// UpdateRoom implements rooms.DataRoomService.
func (r *RoomService) UpdateRoom(roomid uint, userid uint, room rooms.Room) error {
	if roomid == 0 {
		return errors.New("invalid room ID")
	}

	if room.RoomName == "" || room.Location == "" || room.Description == "" || room.Price == 0 || room.QuantityBedroom == 0 || room.QuantityGuest == 0 {
		return errors.New("[validation] roomname/location/description/roomprice/quantitybedroom/quantityguest tidak boleh kosong")
	}

	cekuserid, err := r.roomData.SelectByUserID(roomid)
	if err != nil {
		return err
	}

	if cekuserid.UserID != userid {
		return errors.New("user id tidak sesuai")
	}

	return r.roomData.UpdateRoom(roomid, room)

}

// GetAllRooms implements rooms.DataRoomService.
func (r *RoomService) GetAllRooms(roomName string) ([]*rooms.Room, error) {
	if roomName != "" {
		// Jika nama ruangan diberikan, lakukan pencarian berdasarkan nama ruangan
		room, err := r.roomData.GetRoomByName(roomName)
		if err != nil {
			return nil, err
		}
		return []*rooms.Room{room}, nil
	}

	// Jika tidak ada nama ruangan yang diberikan, kembalikan semua ruangan
	allRooms, err := r.roomData.GetAllRooms()
	if err != nil {
		return nil, err
	}

	// Konversi slice rooms.Room menjadi []*rooms.Room
	var allRoomsPtr []*rooms.Room
	for i := range allRooms {
		allRoomsPtr = append(allRoomsPtr, &allRooms[i])
	}

	return allRoomsPtr, nil
}

// GetRoomByID implements rooms.DataRoomService.
func (r *RoomService) GetRoomByID(roomID uint) (*rooms.Room, error) {
	if roomID <= 0 {
		return nil, errors.New("id not valid")
	}
	return r.roomData.GetRoomByID(roomID)
}

func (r *RoomService) GetUserRooms(userID uint) (*rooms.Room, error) {
	rooms, err := r.roomData.SelectByUserID(userID)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
