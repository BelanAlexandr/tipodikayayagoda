package repository

import (
	"fmt"
)

func Register(login, password, role string) error {
	_, err := db.Exec(
		"INSERT INTO users(login,pass,role) VALUES($1, $2,$3)",
		login,
		password,
		role,
	)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}

	return nil
}
