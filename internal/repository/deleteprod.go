package repository

func DeleteProd(id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
