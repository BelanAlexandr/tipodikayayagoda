package service

import (
	"fmt"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProducts(role int, userID int, search string, page int, limit int, sort string, category int) ([]models.Product, int, error) {

	offset := (page - 1) * limit

	if role == models.RoleAdmin {

		products, totalCount, err := repository.GetAllProdAdmin(search, limit, offset, sort, category)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to get admin products: %w", err)
		}

		return products, totalCount, nil
	} else if role == models.RoleClient {

		products, totalCount, err := repository.GetAllProdClient(search, limit, offset, sort, category)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to get client products: %w", err)
		}

		return products, totalCount, nil
	}
	products, totalCount := repository.GetProdpoID(userID, search, limit, offset, sort, category)

	return products, totalCount, nil
}
