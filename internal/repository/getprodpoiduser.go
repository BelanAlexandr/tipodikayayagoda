package repository

import "tipodikayayagoda/internal/models"

func GetProdpoID(userID int) []models.Product {

	rows, err := db.Query(`
			SELECT id, name, description, price, count, seller_id
			FROM products
			WHERE seller_id = $1
		`, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count, &product.SellerID)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}
