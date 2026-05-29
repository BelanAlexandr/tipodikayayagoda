package repository

import (
	"fmt"
	"tipodikayayagoda/internal/models"
)

func Updateproduct(product models.Product) error {

	productQuery := `
			UPDATE products 
			SET name = $1, description = $2, category_id = $3, img_url = $4 
			WHERE id = $5;`

	_, err := db.Exec(productQuery, product.Name, product.Description, product.Category_id, product.ImgURL, product.ID)
	if err != nil {
		return fmt.Errorf("failed to update products table: %w", err)
	}

	return nil
}
