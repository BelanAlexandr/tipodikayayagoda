package repository

func MarkNotificationAsRead(notifID, userID int) error {

	query := "UPDATE notifications SET is_read = true WHERE id = $1 AND user_id = $2"
	_, err := db.Exec(query, notifID, userID)
	return err
}
