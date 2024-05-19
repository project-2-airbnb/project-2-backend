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
	var userData Users
	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var users = users.User{
		UserID:      userData.ID,
		UserPicture: userData.UserPicture,
		UserName:    userData.UserName,
		Email:       userData.Email,
		Password:    userData.Password,
		Phone:       userData.Phone,
		Address:     userData.Address,
		UserType:    userData.UserType,
	}

	return &users, nil
}

// AccountById implements users.DataUserInterface.
func (u *userQuery) AccountById(userid uint) (*users.User, error) {
	var userData Users
	tx := u.db.First(&userData, userid)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var users = users.User{
		UserID:      userData.ID,
		UserPicture: userData.UserPicture,
		UserName:    userData.UserName,
		Email:       userData.Email,
		Password:    userData.Password,
		Phone:       userData.Phone,
		Address:     userData.Address,
		UserType:    userData.UserType,
	}

	return &users, nil
}

// CreateAccount implements users.DataUserInterface.
func (u *userQuery) CreateAccount(account users.User) error {
	userGorm := Users{
		UserPicture: account.UserPicture,
		UserName:    account.UserName,
		Email:       account.Email,
		Password:    account.Password,
		Phone:       account.Phone,
		Address:     account.Address,
		UserType:    account.UserType,
	}
	tx := u.db.Create(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil

}

// DeleteAccount implements users.DataUserInterface.
func (u *userQuery) DeleteAccount(userid uint) error {
	tx := u.db.Delete(&Users{}, userid)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateAccount implements users.DataUserInterface.
func (u *userQuery) UpdateAccount(userid uint, account users.User) error {
	var userGorm Users
	tx := u.db.First(&userGorm, userid)
	if tx.Error != nil {
		return tx.Error
	}

	userGorm.UserPicture = account.UserPicture
	userGorm.UserName = account.UserName
	userGorm.Email = account.Email
	userGorm.Password = account.Password
	userGorm.Phone = account.Phone
	userGorm.Address = account.Address

	tx = u.db.Save(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}
