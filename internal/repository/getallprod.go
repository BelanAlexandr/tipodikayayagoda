package repository

import (
	"database/sql"
	"tipodikayayagoda/internal/models"
)

func GetAllProd() []models.Product {
	var imgURL sql.NullString
	var desc sql.NullString
	res, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}
	var products []models.Product
	for res.Next() {
		var product models.Product
		err := res.Scan(&product.ID, &product.Name, &desc, &product.Price, &product.Count, &product.SellerID, &imgURL)
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
	return products
}
