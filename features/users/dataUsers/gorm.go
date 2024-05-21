package datausers

import (
	datareservations "project-2/features/reservations/dataReservations"
	datareview "project-2/features/review/dataReview"
	datarooms "project-2/features/rooms/dataRooms"
	"project-2/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName       string
	Email          string `gorm:"unique"`
	Password       string
	RetypePassword string
	Address        string
	PhoneNumber    string
	PictureProfile string
	UserType       string
	Rooms          []datarooms.Rooms               `gorm:"foreignKey:UserID"`
	Reservations   []datareservations.Reservations `gorm:"foreignKey:UserID"`
	Reviews        []datareview.Reviews            `gorm:"foreignKey:UserID"`
}

func (u Users) IsValidRole() bool {
	return u.UserType == "customer" || u.UserType == "hosting"
}

func (u Users) ModelToUser() users.User {
	return users.User{
		UserID:         u.ID,
		FullName:       u.FullName,
		Email:          u.Email,
		Address:        u.Address,
		PhoneNumber:    u.PhoneNumber,
		PictureProfile: u.PictureProfile,
		UserType:       u.UserType,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}
