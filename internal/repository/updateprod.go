package repository

import (
	"fmt"
	"tipodikayayagoda/internal/models"
)

func Updateproduct(product models.Product, offer models.ProductOffer, role int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start update tx: %w", err)
	}
	defer tx.Rollback()
	if role == models.RoleAdmin {
		productQuery := `
			UPDATE products 
			SET name = $1, description = $2, category_id = $3, img_url = $4 
			WHERE id = $5;`

		_, err = tx.Exec(productQuery, product.Name, product.Description, product.Category_id, product.ImgURL, product.ID)
		if err != nil {
			return fmt.Errorf("failed to update products table: %w", err)
		}
	}
	offerQuery := `
		INSERT INTO product_offers (product_id, seller_id, price, count) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (product_id, seller_id) 
		DO UPDATE SET price = EXCLUDED.price, count = EXCLUDED.count;`

	_, err = tx.Exec(offerQuery, offer.ProductID, offer.SellerID, offer.Price, offer.Count)
	if err != nil {
		return fmt.Errorf("failed to update product_offers table: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit update tx: %w", err)
	}

	return nil
}
