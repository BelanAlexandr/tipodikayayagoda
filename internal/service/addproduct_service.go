package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func Addproduct(product models.Product, role string, userID int) error {
	if role == "admin" {
		return repository.Addproduct(product)
	} else if role == "seller" {
		product.SellerID = userID
		return repository.Addproduct(product)
	}
	return nil
}
