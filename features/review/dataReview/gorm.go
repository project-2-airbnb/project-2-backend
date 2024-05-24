package datareview

import (
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	UserID  uint   
	RoomID  uint   
	Rating  int    
	Comment string 
}
