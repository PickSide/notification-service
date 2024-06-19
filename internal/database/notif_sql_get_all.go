package database

import (
	"errors"
	"notification-service/pkg/models"
	"notification-service/pkg/utils"
)

func GetUserNotifications(userID string) ([]models.Notification, error) {
	uint64UserID, err := utils.StringToUint64(userID)
	if err != nil {
		return []models.Notification{}, errors.New("GetUserNotifications - Error during conversion (Malformed ID)")
	}

	rows, err := GetClient().Query("SELECT id, expires, extra, is_read, type FROM notifications WHERE recipient_id = ?", uint64UserID)
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
