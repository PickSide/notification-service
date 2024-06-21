package database

func UpdateSeenStatus(notificationID string) error {
	tx, err := GetClient().Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()
	_, err = tx.Exec(
		`UPDATE notifications SET is_read = 1 WHERE id = ?`,
		notificationID,
	)

	return err
}
