package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProducts(role string, userID int) ([]models.Product, error) {
	if role == "admin" {
		return repository.GetAllProd(), nil
	}
	return repository.GetProdpoID(role, userID), nil
}
