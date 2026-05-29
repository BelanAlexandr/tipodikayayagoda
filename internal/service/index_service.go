package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProducts(role int, userID int, search string, page int, limit int, sort string, category int) ([]models.Product, int, error) {

	offset := (page - 1) * limit

	if role == models.RoleAdmin {

		products, totalCount := repository.GetAllProdAdmin(search, limit, offset, sort, category)

		return products, totalCount, nil
	} else if role == models.RoleClient {

		products, totalCount := repository.GetAllProdClient(search, limit, offset, sort, category)

		return products, totalCount, nil
	}
	products, totalCount := repository.GetProdpoID(userID, search, limit, offset, sort, category)

	return products, totalCount, nil
}
