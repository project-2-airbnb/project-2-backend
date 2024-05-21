package handler

type UserResponse struct {
	FullName       string `json:"fullname"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	PhoneNumber    string `json:"phone_number"`
	PictureProfile string `json:"picture_profile"`
}
