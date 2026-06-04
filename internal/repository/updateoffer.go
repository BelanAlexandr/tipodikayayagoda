package repository

import "fmt"

func UpdateOffer(id int, price float64, count int, user_id int) error {

	productQuery := `
			UPDATE product_offers 
			SET price = $1, count = $2 
			WHERE product_id = $3 AND seller_id=$4;`

	_, err := db.Exec(productQuery, price, count, id, user_id)
	productQuery = `
        UPDATE products 
        SET min_price = $2  
        WHERE id = $1 
          AND (min_price > $2 OR min_price = 0.00);`
	_, err = db.Exec(productQuery, id, price)
	if err != nil {
		return fmt.Errorf("failed to update products table: %w", err)
	}

	return nil
}
