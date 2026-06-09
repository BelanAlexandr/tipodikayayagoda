package repository

import (
	"fmt"
	"time"
	"tipodikayayagoda/internal/models"
)

func Register(user models.User, createdAt time.Time) (int, error) {

	var exists bool
	err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE login=$1)",
		user.Login,
	).Scan(&exists)

	if exists {
		return -1, fmt.Errorf("user with login %s already exists", user.Login)
	}
	_, err = db.Exec(
		"INSERT INTO users(login,pass,name,secondname,role,date) VALUES($1, $2,$3, $4, $5,$6)",
		user.Login,
		user.Password,
		user.Name,
		user.SecondName,
		user.Role,
		createdAt,
	)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return -1, err
	}
	var id int
	err = db.QueryRow(
		"SELECT id FROM users WHERE login=$1",
		user.Login,
	).Scan(&id)
	if err != nil {
		fmt.Println("Error checking user existence:", err)
		return -1, err
	}
	return id, nil
}
