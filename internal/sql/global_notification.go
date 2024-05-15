package sqlutils

import (
	"notification-service/internal/database"
	"notification-service/pkg/models"
	"time"
)

func CreateGlobalNotification(req models.CreateGlobalNotificationRequest) (int64, error) {
	tx, err := database.GetClient().Begin()
	if err != nil {
		return -1, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	result, err := tx.Exec(
		`INSERT INTO global_notifications (content, expires, title) VALUES (?, ?, ?)`,
		req.Content, time.Now().AddDate(0, 0, 4), req.Title,
	)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
