package service

import (
	"fmt"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
)

func GetProducts(role int, userID int, search string, lastID int, lastPrice float64, lastRank float64, lastLength int, limit int, sort string, category int) ([]models.Product, int, error) {

	if role == models.RoleAdmin {

		products, totalCount, err := repository.GetAllProdAdmin(search, limit, lastID, lastPrice, lastRank, lastLength, sort, category)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to get admin products: %w", err)
		}
		return products, totalCount, nil

	} else if role == models.RoleClient {

		products, totalCount, err := repository.GetAllProdClient(search, limit, lastID, lastPrice, lastRank, lastLength, sort, category)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to get client products: %w", err)
		}
		return products, totalCount, nil
	}

	products, totalCount, err := repository.GetProdpoID(userID, search, limit, lastID, lastPrice, lastRank, lastLength, sort, category)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get products by ID: %w", err)
	}

	return products, totalCount, nil
}
