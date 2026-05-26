package repository

import "tipodikayayagoda/internal/models"

func Updateproduct(product models.Product) error {

	_, err := db.Exec("update products set name=$1, description=$2, price=$3, count=$4, seller_id=$5, img_url=$6, category_id=$7 where id=$8",
		product.Name, product.Description, product.Price, product.Count, product.SellerID, product.ImageURL, product.Category_id, product.ID)
	return err
}
