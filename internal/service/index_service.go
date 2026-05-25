package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProducts(role int, userID int) ([]models.Product, error) {
	if role == models.RoleAdmin {
		return repository.GetAllProd(), nil
	}
	if role == models.RoleClient {
		return repository.GetAllProd(), nil
	}
	return repository.GetProdpoID(userID), nil
}
