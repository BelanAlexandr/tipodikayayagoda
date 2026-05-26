package repository

import "tipodikayayagoda/internal/models"

func Addproduct(product models.Product) error {
	_, err := db.Exec("INSERT INTO products (name, description, price,count,seller_id,category_id) VALUES ($1, $2, $3, $4, $5,$6)", product.Name, product.Description, product.Price, product.Count, product.SellerID, product.Category_id)
	if err != nil {
		return err
	}
	return nil
}
