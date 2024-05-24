package handler

type ReviewRequest struct {
	RoomID  uint   `json:"room_id" form:"room_id"`
	Rating  int    `json:"rating" form:"rating"`
	Comment string `json:"comment" form:"comment"`
}