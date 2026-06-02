package repository

import (
	"database/sql"
	"tipodikayayagoda/internal/models"
)

func AllProd(searchQuery string) ([]models.Product, error) {
	query := "SELECT id, name, description, img_url, offer, category_id FROM products WHERE name ILIKE $1 LIMIT 15"

	rows, err := db.Query(query, "%"+searchQuery+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		var desc sql.NullString
		var imgURL sql.NullString

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&desc,
			&imgURL,
			&product.Offer,
			&product.Category_id,
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
