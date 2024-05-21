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

// CreateAccount implements users.DataUserInterface.
func (u *userQuery) CreateAccount(account users.User) error {
	userGorm := Users{
		FullName:       account.FullName,
		Email:          account.Email,
		Password:       account.Password,
		RetypePassword: account.RetypePassword,
		Address:        account.Address,
		PhoneNumber:    account.PhoneNumber,
		UserType:       account.UserType,
		PictureProfile: account.PictureProfile,
	}
	tx := u.db.Create(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteAccount implements users.DataUserInterface.
func (u *userQuery) DeleteAccount(userid uint) error {
	panic("not implemented")
}

// UpdateAccount implements users.DataUserInterface.
func (u *userQuery) UpdateAccount(userid uint, account users.User) error {
	panic("not implemented")
}

// AccountByEmail implements users.DataUserInterface.
func (u *userQuery) AccountByEmail(email string) (*users.User, error) {
	panic("not implemented")
}

// AccountById implements users.DataUserInterface.
func (u *userQuery) AccountById(userid uint) (*users.User, error) {
	panic("not implemented")
}
