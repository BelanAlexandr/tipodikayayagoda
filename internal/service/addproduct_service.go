package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func Addproduct(product models.Product, role int, userID int) error {
	if role == models.RoleAdmin {
		if product.SellerID == 0 {
			return errors.New("seller ID is required for admin")
		}
		if !repository.SellerCheck(product.SellerID) {
			return errors.New("invalid seller ID")
		}
		return repository.Addproduct(product)
	} else if role == models.RoleSeller {
		product.SellerID = userID
		return repository.Addproduct(product)
	}
	return errors.New("invalid user role")
}
