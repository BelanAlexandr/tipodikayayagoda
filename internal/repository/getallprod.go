package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func GetAllProdAdmin(search string, limit int, lastID int, lastPrice float64, sort string, categoryID int) ([]models.Product, int, error) {
	return getAllProducts(search, limit, lastID, lastPrice, sort, categoryID, false)
}

func GetAllProdClient(search string, limit int, lastID int, lastPrice float64, sort string, categoryID int) ([]models.Product, int, error) {
	return getAllProducts(search, limit, lastID, lastPrice, sort, categoryID, true)
}

func getAllProducts(search string, limit int, lastID int, lastPrice float64, sort string, categoryID int, clientOnly bool) ([]models.Product, int, error) {
	var totalCount int

	// ГИБРИДНЫЙ ПОИСК:
	// 1. Ищет точный корень слова через быстрый FTS индекс: to_tsvector @@ plainto_tsquery
	// 2. Ищет вхождение подстроки через ILIKE (чтобы успешно находить "Товар№4" и "Товар №4")
	whereClause := `WHERE ($1 = '' OR 
		to_tsvector('russian', p.name) @@ plainto_tsquery('russian', $1) OR 
		p.name ILIKE '%' || $1 || '%'
	) AND ($2 = 0 OR p.category_id = $2)`

	if clientOnly {
		whereClause += " AND (p.offer = true)"
	}

	var countQuery string
	if clientOnly {
		countQuery = `
			SELECT COUNT(DISTINCT p.id) 
			FROM products p
			INNER JOIN product_offers o ON p.id = o.product_id
		` + whereClause
	} else {
		countQuery = `
			SELECT COUNT(*) 
			FROM products p
		` + whereClause
	}

	err := db.QueryRow(countQuery, search, categoryID).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("count query failed: %w", err)
	}

	if totalCount == 0 {
		return []models.Product{}, 0, nil
	}

	cursorClause := ""
	var orderBy string

	switch sort {
	case "price_asc":
		orderBy = "min_price ASC, p.id ASC"
		if lastID > 0 {
			cursorClause = fmt.Sprintf("HAVING (COALESCE(MIN(o.price), 0), p.id) > (%f, %d)", lastPrice, lastID)
		}
	case "price_desc":
		orderBy = "min_price DESC, p.id DESC"
		if lastID > 0 {
			cursorClause = fmt.Sprintf("HAVING (COALESCE(MIN(o.price), 0), p.id) < (%f, %d)", lastPrice, lastID)
		}
	case "id_asc":
		orderBy = "p.id ASC"
		if lastID > 0 {
			whereClause += fmt.Sprintf(" AND p.id > %d", lastID)
		}
	case "new":
		// Если передан поисковый запрос, сначала выводим точные совпадения по длине строки,
		// чтобы "Товар №4" был выше, чем "Товар №4005", а внутри совпадений — сортируем по новизне.
		if search != "" {
			orderBy = "ts_rank(to_tsvector('russian', p.name), plainto_tsquery('russian', $1)) DESC, LENGTH(p.name) ASC, p.id DESC"
		} else {
			orderBy = "p.id DESC"
		}
		if lastID > 0 {
			whereClause += fmt.Sprintf(" AND p.id < %d", lastID)
		}
	default:
		if search != "" {
			orderBy = "ts_rank(to_tsvector('russian', p.name), plainto_tsquery('russian', $1)) DESC, LENGTH(p.name) ASC, p.id DESC"
		} else {
			orderBy = "p.id DESC"
		}
		if lastID > 0 {
			whereClause += fmt.Sprintf(" AND p.id < %d", lastID)
		}
	}

	dataQuery := `
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
	` + whereClause + `
		GROUP BY p.id
	` + cursorClause + `
		ORDER BY ` + orderBy + `
		LIMIT $3
	`

	rows, err := db.Query(dataQuery, search, categoryID, limit)
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
