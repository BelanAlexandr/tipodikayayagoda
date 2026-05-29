package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func AllProd() []models.Product {
	return repository.AllProd()
}
