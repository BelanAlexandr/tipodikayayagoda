package repository

import (
	"log"
)

func AddNotification(user_id int, message string) error {

	querySave := `
    INSERT INTO notifications (user_id, text, is_read, created_at) 
    VALUES ($1, $2, false, NOW())`

	_, err := db.Exec(querySave, user_id, message)
	if err != nil {
		log.Println("Ошибка сохранения уведомления в БД:", err)
	}
	return err
}
