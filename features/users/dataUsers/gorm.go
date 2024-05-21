package datausers

import (
	datareservations "project-2/features/reservations/dataReservations"
	datareview "project-2/features/review/dataReview"
	datarooms "project-2/features/rooms/dataRooms"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName       string                          `json:"fullname"`
	Email          string                          `json:"email"`
	Password       string                          `json:"password"`
	RetypePassword string                          `json:"retype_password"`
	Address        string                          `json:"address"`
	PhoneNumber    string                          `json:"phone_number"`
	PictureProfile string                          `json:"picture_profile"`
	UserType       string                          `json:"user_type"`
	Rooms          []datarooms.Rooms               `gorm:"foreignKey:UserID"`
	Reservations   []datareservations.Reservations `gorm:"foreignKey:UserID"`
	Reviews        []datareview.Reviews            `gorm:"foreignKey:UserID"`
}

func (u Users) IsValidRole() bool {
	return u.UserType == "customer" || u.UserType == "hosting"
}
