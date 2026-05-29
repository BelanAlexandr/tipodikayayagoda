package repository

import "log"

func AddOffer(id, count int, price float64, user_id int) error {
	productQuery := `
        INSERT INTO product_offers (product_id, seller_id, price, count) 
        VALUES ($1, $2, $3, $4) `

	_, err := db.Exec(productQuery, id, user_id, price, count)
	if err != nil {
		log.Println("Ошибка сохранения уведомления в БД:", err)
		return err
	}
	return nil
}
