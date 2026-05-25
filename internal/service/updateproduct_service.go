package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func UpdateProd(product models.Product, userID, role int) error {

	if role == models.RoleAdmin {
		if product.SellerID == 0 {
			return errors.New("seller ID is required for admin")
		}
		if !repository.SellerCheck(product.SellerID) {
			return errors.New("invalid seller ID")
		}
		return repository.Updateproduct(product)
	} else if role == models.RoleSeller {
		product.SellerID = userID
		return repository.Updateproduct(product)
	}
	return errors.New("invalid user role")
}
