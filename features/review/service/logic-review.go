package service

import (
	"errors"
	"project-2/features/review"
)

type ReviewService struct {
	reviewData review.ReviewModel
}

func New(rm review.ReviewModel) review.ReviewService {
	return &ReviewService{
		reviewData: rm,
	}
}

func (r *ReviewService) AddReview(review review.Review) error {
	if review.Comment == "" || review.Rating == 0 || review.RoomID == 0 {
		return errors.New("[validation] comment/rating/roomid tidak boleh kosong")
	}

	err := r.reviewData.AddReview(review)
	if err != nil {
		return err
	}
	return nil
}

