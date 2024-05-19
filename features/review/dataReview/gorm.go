package datareview

import (
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	RoomID  uint   `json:"room_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
