package repository

import (
	"fmt"
	"time"
	"tipodikayayagoda/internal/models"
)

func Register(user models.User, createdAt time.Time) error {

	var exists bool
	err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE login=$1)",
		user.Login,
	).Scan(&exists)

	if exists {
		return fmt.Errorf("user with login %s already exists", user.Login)
	}

	_, err = db.Exec(
		"INSERT INTO users(login,pass,name,role,date) VALUES($1, $2,$3, $4, $5)",
		user.Login,
		user.Password,
		user.Name,
		user.Role,
		createdAt,
	)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}

	return nil
}
