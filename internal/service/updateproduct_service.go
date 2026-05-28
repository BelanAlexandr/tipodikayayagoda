package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func UpdateProd(productID int, name, desc string, categoryID int, price float64, count int, reqSellerID int, userID int, role int) error {
	existingProduct := repository.GetProductpoIID(productID)
	if existingProduct.ID == 0 {
		return errors.New("product not found")
	}
	var finalSellerID int
	if role == models.RoleAdmin {
		if reqSellerID == 0 {
			return errors.New("seller ID is required for admin")
		}
		if !repository.SellerCheck(reqSellerID) {
			return errors.New("invalid seller ID")
		}
		finalSellerID = reqSellerID
	} else if role == models.RoleSeller {
		finalSellerID = userID
	} else {
		return errors.New("invalid user role")
	}

	prod := models.Product{
		ID:          productID,
		Name:        name,
		Description: desc,
		Category_id: categoryID,
		ImgURL:      existingProduct.ImageURL,
	}

	offer := models.ProductOffer{
		ProductID: productID,
		SellerID:  finalSellerID,
		Price:     price,
		Count:     count,
	}

	return repository.Updateproduct(prod, offer, role)
}
