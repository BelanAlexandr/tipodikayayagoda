package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func GetAllProdAdmin(search string, limit int, offset int, sort string, categoryID int) ([]models.Product, int) {
	var imgURL sql.NullString
	var desc sql.NullString
	var totalCount int

	countQuery := `
        SELECT COUNT(*) 
        FROM products p
        WHERE ($1 = '' OR to_tsvector('russian', p.name || ' ' || COALESCE(p.description, '')) @@ plainto_tsquery('russian', $1))
          AND ($2 = 0 OR p.category_id = $2)
    `
	err := db.QueryRow(countQuery, search, categoryID).Scan(&totalCount)
	if err != nil {
		panic(err)
	}

	orderBy := "p.id DESC"
	switch sort {
	case "price_asc":
		orderBy = "min_price ASC"
	case "price_desc":
		orderBy = "min_price DESC"
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
        WHERE ($1 = '' OR to_tsvector('russian', p.name || ' ' || COALESCE(p.description, '')) @@ plainto_tsquery('russian', $1))
          AND ($2 = 0 OR p.category_id = $2)
        GROUP BY p.id
        ORDER BY %s
        LIMIT $3 OFFSET $4
    `, orderBy)

	rows, err := db.Query(dataQuery, search, categoryID, limit, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product

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
			panic(err)
		}
		if desc.Valid {
			product.Description = desc.String
		}
		if imgURL.Valid {
			product.ImgURL = imgURL.String
		}

		products = append(products, product)
	}

	return products, totalCount
}
func GetAllProdClient(search string, limit int, offset int, sort string, categoryID int) ([]models.Product, int) {
	var imgURL sql.NullString
	var desc sql.NullString
	var totalCount int

	countQuery := `
        SELECT COUNT(*) 
        FROM products p
        WHERE ($1 = '' OR to_tsvector('russian', p.name || ' ' || COALESCE(p.description, '')) @@ plainto_tsquery('russian', $1))
          AND ($2 = 0 OR p.category_id = $2) AND (offer=true)
    `
	err := db.QueryRow(countQuery, search, categoryID).Scan(&totalCount)
	if err != nil {
		panic(err)
	}

	orderBy := "p.id DESC"
	switch sort {
	case "price_asc":
		orderBy = "min_price ASC"
	case "price_desc":
		orderBy = "min_price DESC"
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
        WHERE ($1 = '' OR to_tsvector('russian', p.name || ' ' || COALESCE(p.description, '')) @@ plainto_tsquery('russian', $1))
          AND ($2 = 0 OR p.category_id = $2) AND (offer=true)
        GROUP BY p.id
        ORDER BY %s
        LIMIT $3 OFFSET $4
    `, orderBy)

	rows, err := db.Query(dataQuery, search, categoryID, limit, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product

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
			panic(err)
		}
		if desc.Valid {
			product.Description = desc.String
		}
		if imgURL.Valid {
			product.ImgURL = imgURL.String
		}

		products = append(products, product)
	}

	return products, totalCount
}
