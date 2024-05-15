package sqlutils

import (
	"errors"
	"notification-service/internal/database"
	"notification-service/pkg/models"
	"notification-service/pkg/utils"
	"time"
)

func CreateNotification(req models.CreateUserNotificationStruct) (*string, error) {
	tx, err := database.GetClient().Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	uint64RecipientID, err := utils.StringToUint64(req.RecipientID)
	if err != nil {
		return nil, errors.New("CreateNotification - Error during conversion (Malformed ID)")
	}

	result, err := tx.Exec(`INSERT INTO notifications (expires, extra, recipient_id, type) VALUES (?, ?, ?, ?)`,
		time.Now().AddDate(0, 0, 4), req.Extra, uint64RecipientID, req.Type,
	)
	if err != nil {
		return nil, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	lastInsertedIDString := utils.Int64ToString(lastInsertedID)

	return &lastInsertedIDString, nil
}
func GetGlobalNotifications() (*[]models.GlobalNotification, error) {
	rows, err := database.GetClient().Query("SELECT content, expires, title FROM global_notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var globalNotifications []models.GlobalNotification

	for rows.Next() {
		var globalNotification models.GlobalNotification

		err := rows.Scan(
			&globalNotification.Content,
			&globalNotification.Expires,
			&globalNotification.Title,
		)
		if err != nil {
			return nil, err
		}

		globalNotification.IDString = utils.Uint64ToString(globalNotification.ID)
		globalNotifications = append(globalNotifications, globalNotification)
	}

	return &globalNotifications, nil
}
func GetUserNotifications(userID string) (*[]models.Notification, error) {
	uint64UserID, err := utils.StringToUint64(userID)
	if err != nil {
		return nil, errors.New("GetUserNotifications - Error during conversion (Malformed ID)")
	}

	rows, err := database.GetClient().Query("SELECT expires, extra, is_read, type FROM notifications WHERE recipient_id = ?", uint64UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification

	for rows.Next() {
		var notification models.Notification

		err := rows.Scan(
			&notification.Expires,
			&notification.Extra,
			&notification.IsRead,
			&notification.Type,
		)
		if err != nil {
			return nil, err
		}

		notification.IDString = utils.Uint64ToString(notification.ID)

		notifications = append(notifications, notification)
	}

	return &notifications, nil
}
