package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProdPoID(id int, role string, userID int) models.Product {

	if role == "seller" {

		p := repository.GetProductpoID(id)
		if p.SellerID == userID {
			return p
		} else {
			return models.Product{}
		}
	}
	return repository.GetProductpoID(id)
}
