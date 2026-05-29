package repository

import "fmt"

func UpdateOffer(id int, price float64, count int, user_id int) error {

	productQuery := `
			UPDATE product_offers 
			SET price = $1, count = $2 
			WHERE product_id = $3 AND seller_id=$4;`

	_, err := db.Exec(productQuery, price, count, id, user_id)
	if err != nil {
		return fmt.Errorf("failed to update products table: %w", err)
	}

	return nil
}
