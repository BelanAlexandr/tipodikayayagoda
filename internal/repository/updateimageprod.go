package repository

func Updateimage(url string, id int) error {
	_, err := db.Exec("update products set img_url=$1 where id=$2", url, id)
	return err
}
