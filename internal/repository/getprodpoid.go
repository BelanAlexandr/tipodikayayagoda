package repository

import (
	"database/sql"
	"tipodikayayagoda/internal/models"
)

func GetProductpoIID(id int) (models.ProductDetails, error) {
	var imgURL sql.NullString
	var desc sql.NullString
	var p models.ProductDetails
	productQuery := `
		SELECT id, name, description, img_url, category_id 
		FROM products 
		WHERE id = $1;`

	err := db.QueryRow(productQuery, id).Scan(&p.ID, &p.Name, &desc, &imgURL, &p.CategoryID)
	if err != nil {
		return models.ProductDetails{}, err
	}
	if desc.Valid {
		p.Description = desc.String
	}
	if imgURL.Valid {
		p.ImageURL = imgURL.String
	}
	offersQuery := `
    SELECT o.seller_id, COALESCE(u.name, 'Продавец') as seller_name, o.price::integer, o.count
    FROM product_offers o
    LEFT JOIN users u ON o.seller_id = u.id
    WHERE o.product_id = $1
    ORDER BY o.price ASC;`
	rows, err := db.Query(offersQuery, id)

	if err != nil {
		return p, err
	}
	defer rows.Close()

	for rows.Next() {
		var offer models.OfferDetail
		err := rows.Scan(&offer.SellerID, &offer.SellerName, &offer.Price, &offer.Count)
		if err != nil {
			continue
		}

		p.Offers = append(p.Offers, offer)
	}

	return p, nil
}
