package service

import (
	"errors"
	"fmt"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func UpdateProd(product models.Product, userID int, role string) error {
	fmt.Println("Updating product:", product)
	if role == "admin" {
		if product.SellerID == 0 {
			return errors.New("seller ID is required for admin")
		}
		if !repository.SellerCheck(product.SellerID) {
			return errors.New("invalid seller ID")
		}
		return repository.Updateproduct(product)
	} else if role == "seller" {
		product.SellerID = userID
		return repository.Updateproduct(product)
	}
	return errors.New("invalid user role")
}
