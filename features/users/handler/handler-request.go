package handler

type UserRequest struct {
	UserPicture string `json:"user_picture"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	UserType    string `json:"user_type"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
