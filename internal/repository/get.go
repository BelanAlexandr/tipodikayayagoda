package repository

func GetUser(login string) (user_id int, user_password string, err error) {
	var id int
	var password string

	err = db.QueryRow(
		"SELECT id,pass FROM users WHERE login=$1",
		login,
	).Scan(&id, &password)

	return id, password, err
}
