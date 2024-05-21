package users

import "time"

type User struct {
	UserID         uint      `json:"user_id"`
	FullName       string    `json:"fullname"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	RetypePassword string    `json:"retype_password"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone_number"`
	PictureProfile string    `json:"picture_profile"`
	UserType       string    `json:"user_type"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt"`
}

func (u User) IsValidRole() bool {
	return u.UserType == "customer" || u.UserType == "hosting"
}

type DataUserInterface interface {
	CreateAccount(account User) error
	AccountByEmail(email string) (*User, error)
	AccountById(userid uint) (*User, error)
	UpdateAccount(userid uint, account User) error
	DeleteAccount(userid uint) error
}

type ServiceUserInterface interface {
	RegistrasiAccount(accounts User) error
	LoginAccount(email string, password string) (data *User, token string, err error)
	GetProfile(userid uint) (data *User, err error)
	UpdateProfile(userid uint, accounts User) error
	DeleteAccount(userid uint) error
}
