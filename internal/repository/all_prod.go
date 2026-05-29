package repository

import (
	"database/sql"
	"tipodikayayagoda/internal/models"
)

func AllProd() []models.Product {

	rows, err := db.Query("SELECT* FROM products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var imgURL sql.NullString
	var desc sql.NullString
	var products []models.Product
	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&desc,
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
	return products
}
