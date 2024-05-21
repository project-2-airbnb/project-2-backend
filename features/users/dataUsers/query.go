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
	var userGorm Users
	tx := u.db.First(&userGorm, userid)
	if tx.Error != nil {
		return tx.Error
	}

	userGorm.FullName = account.FullName
	userGorm.Email = account.Email
	userGorm.Password = account.Password
	userGorm.RetypePassword = account.RetypePassword
	userGorm.Address = account.Address
	userGorm.PhoneNumber = account.PhoneNumber
	userGorm.PictureProfile = account.PictureProfile

	tx = u.db.Save(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// AccountByEmail implements users.DataUserInterface.
func (u *userQuery) AccountByEmail(email string, usertype string) (*users.User, error) {
	var userData Users
	tx := u.db.Where("email = ? AND user_type = ?", email, usertype).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var users = users.User{
		FullName:       userData.FullName,
		Email:          userData.Email,
		Password:       userData.Password,
		RetypePassword: userData.RetypePassword,
		Address:        userData.Address,
		PhoneNumber:    userData.PhoneNumber,
		UserType:       userData.UserType,
		PictureProfile: userData.PictureProfile,
	}

	return &users, nil
}

// AccountById implements users.DataUserInterface.
func (u *userQuery) AccountById(userid uint) (*users.User, error) {
	panic("not implemented")
}
