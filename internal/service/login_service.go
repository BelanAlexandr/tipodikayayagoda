package service

import (
	"fmt"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func Login(login, password string) (string, error) {

	user_id, user_password, err := repository.GetUser(login)
	if err != nil {
		return "", err
	}

	ok := utils.CheckHash(password, user_password)
	if !ok {
		return "", fmt.Errorf("invalid password")
	}
	token, err := utils.GenerateJWT(user_id)
	if err != nil {
		return "", err
	}
	return token, nil
}
