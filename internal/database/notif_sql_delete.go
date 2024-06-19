package database

import "notification-service/pkg/utils"

func DeleteNotification(notificationID string) error {
	uint64NotificationID, err := utils.StringToUint64(notificationID)
	if err != nil {
		return err
	}

	if _, err := GetClient().Exec("DELETE FROM notifications WHERE id = ?", uint64NotificationID); err != nil {
		return err
	}

	return nil
}
