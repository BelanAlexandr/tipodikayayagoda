package repository

import (
	"database/sql"
	"fmt"
	"tipodikayayagoda/internal/models"
)

func AllProd(searchQuery string) ([]models.Product, error) {
	if searchQuery == "" {
		return nil, nil
	}

	query := `
        SELECT 
            id, 
            name, 
            description, 
            img_url, 
            offer, 
            category_id,
            ts_rank(name_tsvector, plainto_tsquery('russian', $1)) AS rank
        FROM products 
        WHERE name_tsvector @@ plainto_tsquery('russian', $1)
        ORDER BY rank DESC, LENGTH(name) ASC, id DESC 
        LIMIT 15`

	rows, err := db.Query(query, searchQuery)
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		var desc sql.NullString
		var imgURL sql.NullString
		var rank float64

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&desc,
			&imgURL,
			&product.Offer,
			&product.Category_id,
			&rank,
		)
		if err != nil {
			return nil, err
		}

		if desc.Valid {
			product.Description = desc.String
		} else {
			product.Description = ""
		}

		if imgURL.Valid {
			product.ImgURL = imgURL.String
		} else {
			product.ImgURL = ""
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
