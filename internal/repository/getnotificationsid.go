package repository

import (
	"tipodikayayagoda/internal/models"
)

func GetNotificationPoId(user_id int) ([]models.Notification, error) {

	queryList := `
    SELECT id, user_id, text, is_read, created_at 
    FROM notifications 
    WHERE user_id = $1 
    ORDER BY id DESC 
    LIMIT 20`

	rows, err := db.Query(queryList, user_id)
	if err != nil {

		return []models.Notification{}, err
	}
	defer rows.Close()

	var notificationsFromDB []models.Notification
	for rows.Next() {
		var n models.Notification
		err := rows.Scan(&n.ID, &n.User_ID, &n.Text, &n.Is_read, &n.Created_at)
		if err != nil {

			return []models.Notification{}, err
		}
		notificationsFromDB = append(notificationsFromDB, n)
	}
	return notificationsFromDB, nil
}
