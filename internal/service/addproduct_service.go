package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func Addproduct(name, desc string, categoryID int) error {
	product := models.Product{
		Name:        name,
		Description: desc,
		Category_id: categoryID,
		ImgURL:      "",
	}
	return repository.Addproduct(product)
}
