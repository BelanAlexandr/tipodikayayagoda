package repository

func GetUserbyID(id int) (role int, err error) {

	var rol int
	err = db.QueryRow(
		"SELECT role FROM users WHERE id=$1",
		id,
	).Scan(&rol)

	return rol, err
}
