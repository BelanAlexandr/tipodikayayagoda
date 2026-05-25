package service

import (
	"time"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func Register(user models.User, userrole int) error {
	hash, err := utils.HashPassword(user.Password)
	user.Password = hash
	ti := time.Now()
	if err != nil {
		return err
	}
	if userrole != models.RoleAdmin {
		user.Role = models.RoleClient
		return repository.Register(user, ti)
	}

	return repository.Register(user, ti)
}
