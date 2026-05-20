package repository

import (
	"tipodikayayagoda/internal/models"
)

func GetUser(login, table string) (models.User, error) {
	var u models.User

	err := db.QueryRow(
		"SELECT id, login, pass, role FROM users WHERE login=$1",
		login,
	).Scan(&u.ID, &u.Login, &u.Password, &u.Role)

	return u, err
}
