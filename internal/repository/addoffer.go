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
        UPDATE products 
        SET min_price = $2  
        WHERE id = $1 
          AND (min_price > $2 OR min_price = 0.00);`
	_, err = db.Exec(productQuery, id, price)
	if err != nil {
		log.Println("Ошибка сохранения уведомления в БД:", err)
		return err
	}
	return nil
}
