package handler

type UserRequest struct {
	FullName       string `json:"fullname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RetypePassword string `json:"retype_password"`
	Address        string `json:"address"`
	PhoneNumber    string `json:"phone_number"`
	PictureProfile string `json:"picture_profile"`
	UserType       string `json:"user_type"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
