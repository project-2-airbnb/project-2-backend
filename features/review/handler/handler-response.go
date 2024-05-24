package handler

type ReviewResponse struct {
	RoomID  uint   `json:"room_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}