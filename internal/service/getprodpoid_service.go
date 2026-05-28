package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProdPoID(id int, role int, userID int) models.ProductDetails {
	p := repository.GetProductpoIID(id)
	if p.ID == 0 {
		return models.ProductDetails{}
	}
	if role == models.RoleSeller {
		isOwner := false
		for _, offer := range p.Offers {
			if offer.SellerID == userID {
				isOwner = true
				break
			}
		}
		if !isOwner {
			return models.ProductDetails{}
		}
	}
	return p
}
