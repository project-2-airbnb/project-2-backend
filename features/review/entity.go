package review

import "time"

type Review struct {
	ReviewID  uint    	
	UserID    uint    	
	RoomID    uint    	
	Rating    int       
	Comment   string    
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt time.Time 
}

type ReviewModel interface {
	AddReview(review Review) error
}

type ReviewService interface {
	AddReview(review Review) error
}
