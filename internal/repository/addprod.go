package repository

import (
	"fmt"
	"tipodikayayagoda/internal/models"
)

func Addproduct(product models.Product) error {

	productQuery := `
        INSERT INTO products (name, description, img_url, category_id) 
        VALUES ($1, $2, $3, $4) 
        `

	_, err := db.Exec(productQuery, product.Name, product.Description, product.ImgURL, product.Category_id)
	if err != nil {
		return fmt.Errorf("insert into products failed: %w", err)
	}

	return nil
}
