package datausers

import (
	"project-2/features/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataUserInterface {
	return &userQuery{
		db: db,
	}
}

// AccountByEmail implements users.DataUserInterface.
func (u *userQuery) AccountByEmail(email string) (*users.User, error) {
	panic("not implemented")
}

// AccountById implements users.DataUserInterface.
func (u *userQuery) AccountById(userid uint) (*users.User, error) {
	panic("not implemented")
}

// CreateAccount implements users.DataUserInterface.
func (u *userQuery) CreateAccount(account users.User) error {
	panic("not implemented")
}

// DeleteAccount implements users.DataUserInterface.
func (u *userQuery) DeleteAccount(userid uint) error {
	panic("not implemented")
}

// UpdateAccount implements users.DataUserInterface.
func (u *userQuery) UpdateAccount(userid uint, account users.User) error {
	panic("not implemented")
}
