package service

import (
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func Login(login, password string) error {
	user, err := repository.GetUser(login, "users")
	if err != nil {
		return err
	}
	chek := utils.CheckHash(password, user.Password)
	if !chek {
		return err
	} else {
		return nil
	}

	return nil
}
