package repository

import "tipodikayayagoda/internal/models"

func GetProductpoID(id int) models.Product {

	row := db.QueryRow(`
    SELECT id, name, description, price, count, seller_id
    FROM products
    WHERE id = $1
`, id)
	var p models.Product

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Count,
		&p.SellerID,
	)
	if err != nil {
		return models.Product{}
	}

	return p

}
