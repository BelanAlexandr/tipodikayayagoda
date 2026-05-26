package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func GetProdpoID(userID int, search string, limit int, offset int, sort string, categoryID int) ([]models.Product, int) {
	var imgURL sql.NullString
	var desc sql.NullString

	var totalCount int

	countQuery := `
		SELECT COUNT(*) 
		FROM products 
		WHERE seller_id = $1 
		  AND ($2 = '' OR to_tsvector('russian', name || ' ' || COALESCE(description, '')) @@ plainto_tsquery('russian', $2))
		  AND ($3 = 0 OR category_id = $3)
	`
	err := db.QueryRow(countQuery, userID, search, categoryID).Scan(&totalCount)
	if err != nil {
		panic(err)
	}

	orderBy := "id DESC"
	switch sort {
	case "price_asc":
		orderBy = "price ASC"
	case "price_desc":
		orderBy = "price DESC"
	}

	dataQuery := fmt.Sprintf(`
		SELECT id, name, description, price, count, seller_id, img_url, category_id
		FROM products
		WHERE seller_id = $1 
		  AND ($2 = '' OR to_tsvector('russian', name || ' ' || COALESCE(description, '')) @@ plainto_tsquery('russian', $2))
		  AND ($3 = 0 OR category_id = $3)
		ORDER BY %s
		LIMIT $4 OFFSET $5
	`, orderBy)

	rows, err := db.Query(dataQuery, userID, search, categoryID, limit, offset)
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
			&product.SellerID,
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
			product.ImageURL = imgURL.String
		}
		products = append(products, product)
	}

	return products, totalCount
}
