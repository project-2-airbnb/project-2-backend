package users

import "time"

type User struct {
	UserID         uint
	FullName       string
	Email          string
	Password       string
	RetypePassword string
	Address        string
	PhoneNumber    string
	PictureProfile string
	UserType       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func (u User) IsValidRole() bool {
	return u.UserType == "customer" || u.UserType == "hosting"
}

type DataUserInterface interface {
	CreateAccount(account User) error
	AccountByEmail(email string, usertype string) (*User, error)
	AccountById(userid uint) (*User, error)
	UpdateAccount(userid uint, account User) error
	DeleteAccount(userid uint) error
}

type ServiceUserInterface interface {
	RegistrasiAccount(accounts User) error
	LoginAccount(email string, password string, usertype string) (data *User, token string, err error)
	GetProfile(userid uint) (data *User, err error)
	UpdateProfile(userid uint, accounts User) error
	DeleteAccount(userid uint) error
}
