package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func GetProdpoID(userID int, search string, limit int, offset int, sort string, categoryID int) ([]models.ProductCard, int) {
	var imgURL sql.NullString
	var desc sql.NullString
	var totalCount int
	countQuery := `
        SELECT COUNT(*) 
        FROM products p
        INNER JOIN product_offers o ON p.id = o.product_id
        WHERE o.seller_id = $1 
          AND ($2 = '' OR to_tsvector('russian', p.name || ' ' || COALESCE(p.description, '')) @@ plainto_tsquery('russian', $2))
          AND ($3 = 0 OR p.category_id = $3)
    `
	err := db.QueryRow(countQuery, userID, search, categoryID).Scan(&totalCount)
	if err != nil {
		panic(err)
	}

	orderBy := "p.id DESC"
	switch sort {
	case "price_asc":
		orderBy = "o.price ASC"
	case "price_desc":
		orderBy = "o.price DESC"
	}

	dataQuery := fmt.Sprintf(`
        SELECT p.id, p.name, p.description, o.price, o.count, p.img_url, p.category_id
        FROM products p
        INNER JOIN product_offers o ON p.id = o.product_id
        WHERE o.seller_id = $1 
          AND ($2 = '' OR to_tsvector('russian', p.name || ' ' || COALESCE(p.description, '')) @@ plainto_tsquery('russian', $2))
          AND ($3 = 0 OR p.category_id = $3)
        ORDER BY %s
        LIMIT $4 OFFSET $5
    `, orderBy)

	rows, err := db.Query(dataQuery, userID, search, categoryID, limit, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []models.ProductCard
	for rows.Next() {
		var product models.ProductCard

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&desc,
			&product.MinPrice,
			&product.TotalCount,
			&imgURL,
			&product.CategoryID,
		)
		if err != nil {
			panic(err)
		}

		if desc.Valid {
			product.Description = desc.String
		}
		if imgURL.Valid {
			product.ImageURL = imgURL.String
		}
		products = append(products, product)
	}

	return products, totalCount
}
