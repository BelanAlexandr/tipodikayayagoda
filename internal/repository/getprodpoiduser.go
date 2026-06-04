package repository

import (
	"fmt"
	"strings"
	"tipodikayayagoda/internal/models"
)

func GetProdpoID(userID int, search string, limit int, lastID int, lastPrice float64, lastRank float64, lastLength int, sort string, categoryID int) ([]models.Product, int, error) {
	var totalCount int
	var countConditions []string
	var countArgs []interface{}
	argIdx := 1

	if search != "" {
		searchCond := fmt.Sprintf("(p.name_tsvector @@ plainto_tsquery('russian', $%d) OR p.name ILIKE '%%' || $%d || '%%')", argIdx, argIdx)
		countConditions = append(countConditions, searchCond)
		countArgs = append(countArgs, search)
		argIdx++
	}
	if categoryID > 0 {
		countConditions = append(countConditions, fmt.Sprintf("p.category_id = $%d", argIdx))
		countArgs = append(countArgs, categoryID)
		argIdx++
	}
	countConditions = append(countConditions, fmt.Sprintf("o.seller_id = $%d", argIdx))
	countArgs = append(countArgs, userID)
	argIdx++
	whereClause := ""
	if len(countConditions) > 0 {
		whereClause = "WHERE " + strings.Join(countConditions, " AND ")
	}

	countQuery := `
			SELECT COUNT(DISTINCT p.id) 
			FROM products p
			INNER JOIN product_offers o ON p.id = o.product_id 
		` + whereClause
	err := db.QueryRow(countQuery, countArgs...).Scan(&totalCount)

	if err != nil {
		return nil, 0, fmt.Errorf("count query failed: %w", err)
	}
	if totalCount == 0 {
		return []models.Product{}, 0, nil
	}
	var dataArgs []interface{}
	dIdx := 1

	searchArgNum := 0
	if search != "" {
		dataArgs = append(dataArgs, search)
		searchArgNum = dIdx
		dIdx++
	}

	catArgNum := 0
	if categoryID > 0 {
		dataArgs = append(dataArgs, categoryID)
		catArgNum = dIdx
		dIdx++
	}
	var dataConditions []string
	if search != "" {
		searchCond := fmt.Sprintf("(p.name_tsvector @@ plainto_tsquery('russian', $%d) OR p.name ILIKE '%%' || $%d || '%%')", searchArgNum, searchArgNum)
		dataConditions = append(dataConditions, searchCond)
	}
	if categoryID > 0 {
		dataConditions = append(dataConditions, fmt.Sprintf("p.category_id = $%d", catArgNum))
	}
	var orderBy string
	switch sort {
	case "price_asc":
		orderBy = "p.min_price ASC, p.id ASC"
		if lastID > 0 {
			dataConditions = append(dataConditions, fmt.Sprintf("(p.min_price, p.id) > ($%d, $%d)", dIdx, dIdx+1))
			dataArgs = append(dataArgs, lastPrice, lastID)
			dIdx += 2
		}
	case "price_desc":
		orderBy = "p.min_price DESC, p.id DESC"
		if lastID > 0 {
			dataConditions = append(dataConditions, fmt.Sprintf("(p.min_price, p.id) < ($%d, $%d)", dIdx, dIdx+1))
			dataArgs = append(dataArgs, lastPrice, lastID)
			dIdx += 2
		}
	case "id_asc":
		orderBy = "p.id ASC"
		if lastID > 0 {
			dataConditions = append(dataConditions, fmt.Sprintf("p.id > $%d", dIdx))
			dataArgs = append(dataArgs, lastID)
			dIdx++
		}
	default:
		if search != "" {
			orderBy = "rank DESC, LENGTH(p.name) ASC, p.id DESC"
			if lastID > 0 {
				dataConditions = append(dataConditions, fmt.Sprintf(`(
					ts_rank(p.name_tsvector, plainto_tsquery('russian', $%d)) < $%d
					OR (ts_rank(p.name_tsvector, plainto_tsquery('russian', $%d)) = $%d AND LENGTH(p.name) > $%d)
					OR (ts_rank(p.name_tsvector, plainto_tsquery('russian', $%d)) = $%d AND LENGTH(p.name) = $%d AND p.id < $%d)
				)`, searchArgNum, dIdx, searchArgNum, dIdx, dIdx+1, searchArgNum, dIdx, dIdx+1, dIdx+2))
				dataArgs = append(dataArgs, lastRank, lastLength, lastID)
				dIdx += 3
			}
		} else {
			orderBy = "p.id DESC"
			if lastID > 0 {
				dataConditions = append(dataConditions, fmt.Sprintf("p.id < $%d", dIdx))
				dataArgs = append(dataArgs, lastID)
				dIdx++
			}
		}
	}
	return products, totalCount, nil
}
