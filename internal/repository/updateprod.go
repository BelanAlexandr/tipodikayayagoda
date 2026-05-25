package repository

import "tipodikayayagoda/internal/models"

func Updateproduct(product models.Product) error {

	_, err := db.Exec("update products set name=$1, description=$2, price=$3, count=$4, seller_id=$5, img_url=$6 where id=$7",
		product.Name, product.Description, product.Price, product.Count, product.SellerID, product.ImageURL, product.ID)
	return err
}
