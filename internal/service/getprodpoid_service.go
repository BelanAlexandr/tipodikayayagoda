package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProdPoID(id int, role int, userID int) (models.ProductDetails, error) {

	p, err := repository.GetProductpoIID(id)
	if err != nil {
		return models.ProductDetails{}, err
	}

	if p.ID == 0 {
		return models.ProductDetails{}, errors.New("product not found")
	}

	return p, nil
}
