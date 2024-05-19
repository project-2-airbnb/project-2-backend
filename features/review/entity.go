package review

import "time"

type Review struct {
	ReviewID  string    `json:"review_id"`
	UserID    string    `json:"user_id"`
	RoomID    string    `json:"room_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
