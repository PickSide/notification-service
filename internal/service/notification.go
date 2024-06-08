package service

import (
	sqlutils "notification-service/internal/sql"
	"notification-service/pkg/models"
)

func CreateNotification(req models.CreateNotificationStruct) error {
	return sqlutils.CreateNotification(req)
}
func DeleteNotification(notificationID string) error {
	return sqlutils.DeleteNotification(notificationID)
}
func GetUserNotifications(userID string) ([]models.Notification, error) {
	return sqlutils.GetUserNotifications(userID)
}
func UpdateSeenStatus(notificationID string) error {
	return sqlutils.UpdateSeenStatus(notificationID)
}
