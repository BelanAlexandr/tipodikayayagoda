package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func BuyProduct(productID int, role int, count int) error {
	if role != models.RoleClient {
		return errors.New("unauthorized")
	}
	pro := repository.GetProductpoIID(productID)
	if pro.ID == 0 {
		return errors.New("product not found")
	}
	if len(pro.Offers) == 0 {
		return errors.New("product out of stock")
	}
	bestOffer := pro.Offers[0]

	if count <= 0 || count > bestOffer.Count {
		return errors.New("not enough items in stock from this seller")
	}

	// 6. Передаем в репозиторий ID товара, количество И ID конкретного продавца, у которого покупаем
	return repository.BuyProduct(productID, bestOffer.SellerID, count)
}
