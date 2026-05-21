package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProdPoID(id int, role string, userID int) models.Product {
	if role == "admin" {
		return repository.GetProductpoID(id)
	}
}
