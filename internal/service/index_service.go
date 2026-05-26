package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProducts(role int, userID int, search string, page int, limit int, sort string) ([]models.Product, int, error) {

	offset := (page - 1) * limit

	if role == models.RoleAdmin || role == models.RoleClient {

		products, totalCount := repository.GetAllProd(search, limit, offset, sort)

		return products, totalCount, nil
	}
	products, totalCount := repository.GetProdpoID(userID, search, limit, offset, sort)

	return products, totalCount, nil
}
