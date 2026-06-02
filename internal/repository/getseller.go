package repository

import (
	"tipodikayayagoda/internal/models"
)

func GetSellers() ([]models.User, error) {
	query := `
		SELECT id, name, secondname 
		FROM users 
		WHERE role = $1
		ORDER BY secondname ASC, name ASC
	`
	rows, err := db.Query(query, models.RoleSeller)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sellers []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.SecondName)
		if err != nil {
			return nil, err
		}
		sellers = append(sellers, u)
	}

	return sellers, nil
}
