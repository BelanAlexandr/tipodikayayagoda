package repository

func GetCategoryID(id int) (string, error) {
	var categoryName string
	err := db.QueryRow("SELECT name FROM categories WHERE id = $1", id).Scan(&categoryName)
	if err != nil {
		return "", err
	}
	return categoryName, nil
}
