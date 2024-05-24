package datareview

import (
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	UserID  uint    `json:"user_id"`
	RoomID  uint    `json:"room_id"`
	Rating  float32 `json:"rating"`
	Comment string  `json:"comment"`
}
