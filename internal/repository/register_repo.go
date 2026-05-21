package repository

import (
	"fmt"
)

func Register(login, password, role string) error {

	var exists bool
	err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE login=$1)",
		login,
	).Scan(&exists)

	if exists {
		return fmt.Errorf("user with login %s already exists", login)
	}

	_, err = db.Exec(
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
