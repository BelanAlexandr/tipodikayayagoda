package repository

import "tipodikayayagoda/internal/models"

func AddCategory(name string) error {
	_, err := db.Exec("INSERT INTO categories (name) VALUES ($1)", name)
	return err
}

func GetCategories() ([]models.Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}
