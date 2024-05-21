package service

import (
	"project-2/features/users"
	"project-2/utils/encrypts"
)

type userService struct {
	userData    users.DataUserInterface
	hashService encrypts.HashInterface
}

func New(ud users.DataUserInterface, hash encrypts.HashInterface) users.ServiceUserInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
	}

}

// DeleteAccount implements users.ServiceUserInterface.
func (u *userService) DeleteAccount(userid uint) error {
	panic("not implemented")
}

// GetProfile implements users.ServiceUserInterface.
func (u *userService) GetProfile(userid uint) (data *users.User, err error) {
	panic("not implemented")
}

// LoginAccount implements users.ServiceUserInterface.
func (u *userService) LoginAccount(email string, password string) (data *users.User, token string, err error) {
	panic("not implemented")
}

// RegistrasiAccount implements users.ServiceUserInterface.
func (u *userService) RegistrasiAccount(accounts users.User) error {
	panic("not implemented")
}

// UpdateProfile implements users.ServiceUserInterface.
func (u *userService) UpdateProfile(userid uint, accounts users.User) error {
	panic("not implemented")
}
