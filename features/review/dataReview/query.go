package datareview

import (
	"project-2/features/review"

	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) review.ReviewModel {
	return &model{
		db:db,
	}
}

func (m *model) AddReview(review review.Review) error {
	tx := m.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	newReview := Reviews{
		UserID: review.UserID,
		RoomID: review.RoomID,
		Rating: review.Rating,
		Comment: review.Comment,
	}

	if err := tx.Create(&newReview).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
