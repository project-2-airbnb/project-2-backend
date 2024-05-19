package service

import (
	"errors"
	"project-2/app/middlewares"
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
	if userid <= 0 {
		return errors.New("id not valid")
	}
	return u.userData.DeleteAccount(userid)
}

// GetProfile implements users.ServiceUserInterface.
func (u *userService) GetProfile(userid uint) (data *users.User, err error) {
	if userid <= 0 {
		return nil, errors.New("id not valid")
	}
	return u.userData.AccountById(userid)
}

// LoginAccount implements users.ServiceUserInterface.
func (u *userService) LoginAccount(email string, password string) (data *users.User, token string, err error) {
	data, err = u.userData.AccountByEmail(email)
	if err != nil {
		return nil, "", err
	}

	isLoginValid := u.hashService.CheckPasswordHash(data.Password, password)
	// ketika isloginvalid = true, maka login berhasil
	if !isLoginValid {
		return nil, "", errors.New("email atau password tidak sesuai")
	}
	token, errJWT := middlewares.CreateToken(int(data.UserID))
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil

}

// RegistrasiAccount implements users.ServiceUserInterface.
func (u *userService) RegistrasiAccount(accounts users.User) error {
	if accounts.UserName == "" || accounts.Email == "" || accounts.Password == "" || accounts.Phone == "" || accounts.Address == "" {
		return errors.New("nama/email/password/phone/address tidak boleh kosong")
	}

	if accounts.UserType != "customer" && accounts.UserType != "hosting" {
		return errors.New("tipe user tidak valid")
	}
	// proses hash password
	result, errHash := u.hashService.HashPassword(accounts.Password)
	if errHash != nil {
		return errHash
	}
	accounts.Password = result
	return u.userData.CreateAccount(accounts)
}

// UpdateProfile implements users.ServiceUserInterface.
func (u *userService) UpdateProfile(userid uint, accounts users.User) error {
	if accounts.UserName == "" || accounts.Email == "" || accounts.Password == "" || accounts.Phone == "" || accounts.Address == "" {
		return errors.New("nama/email/password/phone/address tidak boleh kosong")
	}
	if accounts.Password != "" {
		// proses hash password
		result, errHash := u.hashService.HashPassword(accounts.Password)
		if errHash != nil {
			return errHash
		}
		accounts.Password = result
	}

	return u.userData.UpdateAccount(userid, accounts)
}
