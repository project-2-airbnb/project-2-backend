package handler

type UserRequest struct {
	FullName       string `json:"fullname" form:"fullname"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	RetypePassword string `json:"retype_password" form:"retype_password"`
	Address        string `json:"address" form:"address"`
	PhoneNumber    string `json:"phone_number" form:"phone_number"`
	PictureProfile string `json:"picture_profile" form:"picture_profile"`
	UserType       string `json:"user_type" form:"user_type"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
