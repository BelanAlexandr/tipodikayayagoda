package repository

import (
	"database/sql"
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
		dataConditions = append(dataConditions, fmt.Sprintf("(p.name_tsvector @@ plainto_tsquery('russian', $%d) OR p.name ILIKE '%%' || $%d || '%%')", searchArgNum, searchArgNum))
	}
	if categoryID > 0 {
		dataConditions = append(dataConditions, fmt.Sprintf("p.category_id = $%d", catArgNum))
	}

	dataConditions = append(dataConditions, fmt.Sprintf("EXISTS (SELECT 1 FROM product_offers WHERE product_id = p.id AND seller_id = $%d)", dIdx))
	dataArgs = append(dataArgs, userID)
	SellerIdArgNum := dIdx
	dIdx++

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

	dataWhere := ""
	if len(dataConditions) > 0 {
		dataWhere = "WHERE " + strings.Join(dataConditions, " AND ")
	}

	dataArgs = append(dataArgs, limit)
	limitArgNum := dIdx

	rankSelect := ", 0.0 as rank"
	if search != "" && sort != "price_asc" && sort != "price_desc" && sort != "id_asc" {
		rankSelect = fmt.Sprintf(", ts_rank(p.name_tsvector, plainto_tsquery('russian', $%d)) as rank", searchArgNum)
	}

	dataQuery := fmt.Sprintf(`
		SELECT 
			p.id, 
			p.name, 
			p.description,  
			COALESCE(agg.price, 0.0) as price,    
			COALESCE(agg.count, 0) as total_count, 
			p.img_url, 
			p.category_id 
			%s                                     
		FROM products p
		LEFT JOIN LATERAL (
			SELECT count, price
			FROM product_offers 
			WHERE seller_id = $%d
			  AND product_id = p.id
			LIMIT 1
		) agg ON true
		%s
		ORDER BY %s
		LIMIT $%d`, rankSelect, SellerIdArgNum, dataWhere, orderBy, limitArgNum)

	rows, err := db.Query(dataQuery, dataArgs...)
	if err != nil {
		fmt.Println(err)
		return nil, 0, fmt.Errorf("data query failed: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var imgURL sql.NullString
		var desc sql.NullString

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&desc,
			&product.Price,
			&product.Count,
			&imgURL,
			&product.Category_id,
			&product.Rank,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("row scan failed: %w", err)
		}

		if desc.Valid {
			product.Description = desc.String
		}
		if imgURL.Valid {
			product.ImgURL = imgURL.String
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows loop failed: %w", err)
	}

	return products, totalCount, nil
}
