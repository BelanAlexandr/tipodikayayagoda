package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func BuyProduct(productID int, Role int, count int) error {
	if Role != models.RoleClient {
		return errors.New("unauthorized")
	}
	pro := repository.GetProductpoIID(productID)
	if pro.ID == 0 {
		return errors.New("product not found")
	}
	if count <= 0 || count > pro.Count {
		return errors.New("invalid count")
	}

	return repository.BuyProduct(productID, count)
}
