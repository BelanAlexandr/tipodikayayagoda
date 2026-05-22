package repository

import (
	"database/sql"
	"tipodikayayagoda/internal/models"
)

func GetProductpoID(id int) models.Product {

	row := db.QueryRow(`
    SELECT id, name, description, price, count, seller_id, img_url
    FROM products
    WHERE id = $1
`, id)
	var imgURL sql.NullString
	var desc sql.NullString
	var p models.Product

	err := row.Scan(
		&p.ID,
		&p.Name,
		&desc,
		&p.Price,
		&p.Count,
		&p.SellerID,
		&imgURL,
	)

	if err != nil {
		return models.Product{}
	}
	if desc.Valid {
		p.Description = desc.String
	}
	if imgURL.Valid {
		p.ImageURL = imgURL.String
	}
	return p

}
