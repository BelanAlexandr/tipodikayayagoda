package repository

import "tipodikayayagoda/internal/models"

func GetAllProd() []models.Product {
	res, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}
	var products []models.Product
	for res.Next() {
		var product models.Product
		err := res.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count, &product.SellerID)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}
