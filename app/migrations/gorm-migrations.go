package migrations

import (
	datareservations "project-2/features/reservations/dataReservations"
	datareview "project-2/features/review/dataReview"
	datarooms "project-2/features/rooms/dataRooms"
	datausers "project-2/features/users/dataUsers"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&datausers.Users{})
	db.AutoMigrate(&datarooms.Rooms{})
	db.AutoMigrate(&datareview.Reviews{})
	db.AutoMigrate(&datareservations.Reservations{})
	db.AutoMigrate(&datarooms.Rooms{})
	db.AutoMigrate(&datarooms.RoomFacilitys{})
	db.AutoMigrate(&datarooms.Facilities{})
}
