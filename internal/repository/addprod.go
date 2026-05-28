package repository

import (
	"fmt"
	"tipodikayayagoda/internal/models"
)

func Addproduct(product models.Product, offer models.ProductOffer) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start tx: %w", err)
	}
	defer tx.Rollback()

	var productID int
	productQuery := `
        INSERT INTO products (name, description, img_url, category_id) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id;`

	err = tx.QueryRow(productQuery, product.Name, product.Description, product.ImgURL, product.Category_id).Scan(&productID)
	if err != nil {
		return fmt.Errorf("insert into products failed: %w", err)
	}

	offerQuery := `
        INSERT INTO product_offers (product_id, seller_id, price, count) 
        VALUES ($1, $2, $3, $4);`

	_, err = tx.Exec(offerQuery, productID, offer.SellerID, offer.Price, offer.Count)
	if err != nil {
		return fmt.Errorf("insert into product_offers failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx: %w", err)
	}

	return nil
}
