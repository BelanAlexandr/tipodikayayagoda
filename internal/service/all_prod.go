package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func AllProd(searchQuery string) ([]models.Product, error) {
	return repository.AllProd(searchQuery)
}
