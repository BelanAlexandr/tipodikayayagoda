package repository

import (
	"fmt"
	"tipodikayayagoda/internal/models"
)

func Addproduct(product models.Product) error {
	var nextID int
	query := `
    SELECT COALESCE(
        (SELECT t1.id + 1 
         FROM users t1 
         LEFT JOIN users t2 ON t1.id + 1 = t2.id 
         WHERE t2.id IS NULL 
         ORDER BY t1.id 
         LIMIT 1), 
        1
    );
`

	err := db.QueryRow(query).Scan(&nextID)
	productQuery := `
        INSERT INTO products (id, name, description, img_url, category_id) 
        VALUES ($1, $2, $3, $4, $5) 
        `

	_, err = db.Exec(productQuery, nextID, product.Name, product.Description, product.ImgURL, product.Category_id)
	if err != nil {
		return fmt.Errorf("insert into products failed: %w", err)
	}

	return nil
}
