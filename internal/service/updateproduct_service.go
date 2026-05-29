package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func UpdateProd(productID int, name, desc string, categoryID int) error {
	existingProduct, err := repository.GetProductpoIID(productID)
	if err != nil {
		return err
	}
	if existingProduct.ID == 0 {
		return errors.New("product not found")
	}
	prod := models.Product{
		ID:          productID,
		Name:        name,
		Description: desc,
		Category_id: categoryID,
		ImgURL:      existingProduct.ImageURL,
	}
	return repository.Updateproduct(prod)
}
