package service

import (
	"fmt"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func Login(login, password string) (string, error) {

	user, err := repository.GetUser(login, "users")
	if err != nil {
		return "", err
	}

	ok := utils.CheckHash(password, user.Password)
	if !ok {
		return "", fmt.Errorf("invalid password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Login, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
