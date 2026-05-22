package service

import (
	"errors"
	"tipodikayayagoda/internal/repository"
)

func BuyProduct(productID int, Role string, count int) error {
	if Role != "client" {
		return errors.New("unauthorized")
	}
	pro := repository.GetProductpoID(productID)
	if pro.ID == 0 {
		return errors.New("product not found")
	}
	if count <= 0 || count > pro.Count {
		return errors.New("invalid count")
	}

	return repository.BuyProduct(productID, count)
}
