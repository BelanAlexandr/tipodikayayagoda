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
	if userrole != models.Roles.AdminID {
		user.Role = models.Roles.ClientID
		return repository.Register(user, ti)
	}
	if user.Role == 1 {
		user.Role = models.Roles.ClientID
	} else if user.Role == 2 {
		user.Role = models.Roles.SellerID
	} else if user.Role == 3 {
		user.Role = models.Roles.AdminID
	}
	return repository.Register(user, ti)
}
