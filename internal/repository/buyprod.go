package repository

import "fmt"

func BuyProduct(productID, sellerID, count int) error {
	query := `
		UPDATE product_offers 
		SET count = count - $1 
		WHERE product_id = $2 AND seller_id = $3 AND count >= $1;
	`
	fmt.Println(productID, sellerID, count)
	_, err := db.Exec(query, count, productID, sellerID)
	if err != nil {
		return err
	}

	return nil
}
