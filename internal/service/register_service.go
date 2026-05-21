package service

import (
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func Register(login, password, role, userrole string) error {
	hash, err := utils.HashPassword(password)

	if err != nil {
		return err
	}
	if userrole != "admin" {
		return repository.Register(login, hash, "client")
	}
	return repository.Register(login, hash, role)
}
