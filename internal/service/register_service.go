package service

import (
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func Register(login, password, role string) error {
	hash, err := utils.HashPassword(password)

	if err != nil {
		return err
	}
	return repository.Register(login, hash, role)
}
