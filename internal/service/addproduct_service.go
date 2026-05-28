package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func Addproduct(name, desc string, categoryID int, price float64, count int, reqSellerID int, role int, userID int) error {
	if role == models.RoleAdmin {
		if reqSellerID == 0 {
			return errors.New("seller ID is required for admin")
		}
		if !repository.SellerCheck(reqSellerID) {
			return errors.New("invalid seller ID")
		}
	} else if role == models.RoleSeller {
		reqSellerID = userID
	} else {
		return errors.New("invalid user role")
	}
	product := models.Product{
		Name:        name,
		Description: desc,
		Category_id: categoryID,
		ImgURL:      "",
	}
	offer := models.ProductOffer{
		SellerID: reqSellerID,
		Price:    price,
		Count:    count,
	}
	return repository.Addproduct(product, offer)
}
