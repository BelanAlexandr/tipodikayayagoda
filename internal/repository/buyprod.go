package repository

import "fmt"

func BuyProduct(productID, count int) error {
	_, err := db.Exec("UPDATE products SET count = count - $2 WHERE id = $1 AND count >= $2", productID, count)
	fmt.Println(err)
	return err
}
