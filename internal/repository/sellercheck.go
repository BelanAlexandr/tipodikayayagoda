package repository

func SellerCheck(sellerID int) bool {
	var exists bool
	_ = db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE id=$1 AND role='27')",
		sellerID,
	).Scan(&exists)

	return exists
}
