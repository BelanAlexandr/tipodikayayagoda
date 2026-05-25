package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProdPoID(id int, role int, userID int) models.Product {

	if role == models.RoleSeller {

		p := repository.GetProductpoID(id)
		if p.SellerID == userID {
			return p
		} else {
			return models.Product{}
		}
	}
	return repository.GetProductpoID(id)
}
