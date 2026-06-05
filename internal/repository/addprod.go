package repository

import (
	"fmt"
	"tipodikayayagoda/internal/models"
)

func Addproduct(product models.Product) error {

	productQuery := `
    INSERT INTO products (id, name, description, img_url, category_id, min_price)
    VALUES (
        COALESCE(
           
            (SELECT t1.id + 1 
             FROM products t1 
         LEFT JOIN products t2 ON t1.id + 1 = t2.id 
         WHERE t2.id IS NULL 
         ORDER BY t1.id 
         LIMIT 1),
        
        1
    ),
    $1, $2, $3, $4,$5
);  `
	_, err := db.Exec(productQuery, product.Name, product.Description, product.ImgURL, product.Category_id, 0.00)
	if err != nil {
		return fmt.Errorf("insert into products failed: %w", err)
	}

	return nil
}
