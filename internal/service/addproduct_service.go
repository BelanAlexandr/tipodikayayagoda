package service

import (
	"errors"
	"fmt"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func Addproduct(product models.Product, role string, userID int) error {
	if role == "admin" {
		if product.SellerID == 0 {
			return errors.New("seller ID is required for admin")
		}
		fmt.Println("Checking seller ID:", product.SellerID)
		if !repository.SellerCheck(product.SellerID) {
			return errors.New("invalid seller ID")
		}
		return repository.Addproduct(product)
	} else if role == "seller" {
		product.SellerID = userID
		return repository.Addproduct(product)
	}
	return errors.New("invalid user role")
}
