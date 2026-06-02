package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func GetAllProdAdmin(search string, limit int, offset int, sort string, categoryID int) ([]models.Product, int, error) {
	return getAllProducts(search, limit, offset, sort, categoryID, false)
}

func GetAllProdClient(search string, limit int, offset int, sort string, categoryID int) ([]models.Product, int, error) {
	return getAllProducts(search, limit, offset, sort, categoryID, true)
}

func getAllProducts(search string, limit int, offset int, sort string, categoryID int, clientOnly bool) ([]models.Product, int, error) {
	var totalCount int

	orderBy := "p.id DESC"
	switch sort {
	case "price_asc":
		orderBy = "min_price ASC"
	case "price_desc":
		orderBy = "min_price DESC"
	}

	extraFilter := ""
	if clientOnly {
		extraFilter = "AND (p.offer = true)"
	}

	if search == "" && categoryID == 0 && !clientOnly {
		err := db.QueryRow("SELECT reltuples::bigint FROM pg_class WHERE relname = 'products'").Scan(&totalCount)
		if err != nil {
			return nil, 0, fmt.Errorf("fast table count failed: %w", err)
		}
	} else {

		countQuery := fmt.Sprintf(`
			SELECT COUNT(*) 
			FROM products p
			WHERE ($1 = '' OR p.name ILIKE '%%' || $1 || '%%')
			  AND ($2 = 0 OR p.category_id = $2) %s
		`, extraFilter)

		err := db.QueryRow(countQuery, search, categoryID).Scan(&totalCount)
		if err != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", err)
		}
	}

	if totalCount == 0 {
		return []models.Product{}, 0, nil
	}

	dataQuery := fmt.Sprintf(`
		SELECT 
			p.id, 
			p.name, 
			p.description, 
			COALESCE(MIN(o.price), 0) as min_price, 
			COALESCE(SUM(o.count), 0) as total_count, 
			p.img_url, 
			p.category_id
		FROM products p
		LEFT JOIN product_offers o ON p.id = o.product_id
		WHERE ($1 = '' OR p.name ILIKE '%%' || $1 || '%%')
		  AND ($2 = 0 OR p.category_id = $2) %s
		GROUP BY p.id
		ORDER BY %s
		LIMIT $3 OFFSET $4
	`, extraFilter, orderBy)

	rows, err := db.Query(dataQuery, search, categoryID, limit, offset)
	if err != nil {
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
