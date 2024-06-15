package sqlutils

import (
	"errors"
	"notification-service/internal/database"
	"notification-service/pkg/models"
	"notification-service/pkg/utils"
	"time"
)

var (
	DAYS_TO_EXPIRE = 4
)

func CreateNotification(req models.CreateNotificationStruct) error {
	uint64RecipientID, err := utils.StringToUint64(req.RecipientID)
	if err != nil {
		return errors.New("CreateNotification - Error during conversion (Malformed ID)")
	}

	tx, err := database.GetClient().Begin()
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

	_, err = tx.Exec(`INSERT INTO notifications (expires, extra, recipient_id, type) VALUES (?, ?, ?, ?)`,
		time.Now().AddDate(0, 0, DAYS_TO_EXPIRE), req.Extra, uint64RecipientID, req.Type,
	)

	return err
}
func DeleteNotification(notificationID string) error {
	uint64NotificationID, err := utils.StringToUint64(notificationID)
	if err != nil {
		return err
	}

	if _, err := database.GetClient().Exec("DELETE FROM notifications WHERE id = ?", uint64NotificationID); err != nil {
		return err
	}

	return nil
}
func GetUserNotifications(userID string) ([]models.Notification, error) {
	uint64UserID, err := utils.StringToUint64(userID)
	if err != nil {
		return []models.Notification{}, errors.New("GetUserNotifications - Error during conversion (Malformed ID)")
	}

	rows, err := database.GetClient().Query("SELECT id, expires, extra, is_read, type FROM notifications WHERE recipient_id = ?", uint64UserID)
	if err != nil {
		return []models.Notification{}, err
	}
	defer rows.Close()

	var notifications []models.Notification

	for rows.Next() {
		var notification models.Notification

		err := rows.Scan(
			&notification.ID,
			&notification.Expires,
			&notification.Extra,
			&notification.IsRead,
			&notification.Type,
		)
		if err != nil {
			return []models.Notification{}, err
		}

		notification.IDString = utils.Uint64ToString(notification.ID)

		notifications = append(notifications, notification)
	}

	return notifications, nil
}
func UpdateSeenStatus(notificationID string) error {
	tx, err := database.GetClient().Begin()
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
