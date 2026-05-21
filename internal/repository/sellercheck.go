package repository

import "fmt"

func SellerCheck(sellerID int) bool {
	var exists bool
	_ = db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE id=$1 AND role='seller')",
		sellerID,
	).Scan(&exists)
	fmt.Println(exists)
	return exists
}
