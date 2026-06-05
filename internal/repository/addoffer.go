package repository

import "log"

func AddOffer(id, count int, price float64, user_id int) error {
	productQuery := `
        INSERT INTO product_offers (product_id, seller_id, price, count) 
        VALUES ($1, $2, $3, $4) `

	_, err := db.Exec(productQuery, id, user_id, price, count)
	productQuery = `UPDATE products 
			SET offer = true
			WHERE id = $1;`

	_, err = db.Exec(productQuery, id)
	productQuery = `
          UPDATE products p
    SET min_price = COALESCE((
        SELECT MIN(price) 
        FROM product_offers 
        WHERE product_id = p.id
    ), 0.00)
    WHERE p.id = $1;`
	_, err = db.Exec(productQuery, id)
	if err != nil {
		log.Println("Ошибка сохранения уведомления в БД:", err)
		return err
	}
	return nil
}
