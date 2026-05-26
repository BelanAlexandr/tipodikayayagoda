package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetSellerId() ([]models.User, error) {

	return repository.GetSellers()
}
