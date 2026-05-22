package repository

func GetImageURL(id int) string {

	row := db.QueryRow(`
    SELECT img_url
    FROM products
    WHERE id = $1
`, id)
	var imgURL string
	err := row.Scan(&imgURL)
	if err != nil {
		return ""
	}
	return imgURL
}
