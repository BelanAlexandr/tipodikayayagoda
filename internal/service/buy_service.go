package service

import (
	"errors"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func BuyProduct(productID int, role int, count int, seller_id int) error {
	if role != models.RoleClient {
		return errors.New("unauthorized")
	}
	pro, err := repository.GetProductpoIID(productID)
	if err != nil {
		return err
	}
	if pro.ID == 0 {
		return errors.New("product not found")
	}
	if len(pro.Offers) == 0 {
		return errors.New("product out of stock")
	}
	var targetOffer *models.OfferDetail // или как у тебя называется структура оффера
	for _, offer := range pro.Offers {
		if offer.SellerID == seller_id { // проверь, как в структуре: SellerID или Seller_id
			targetOffer = &offer
			break
		}
	}
	if count <= 0 || count > targetOffer.Count {
		return errors.New("not enough items in stock from this seller")
	}
	return repository.BuyProduct(productID, seller_id, count)
}
