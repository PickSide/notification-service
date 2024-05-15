package service

import (
	sqlutils "notification-service/internal/sql"
	"notification-service/pkg/models"
)

func CreateGlobalNotification(req models.CreateGlobalNotificationRequest) (int64, error) {
	return sqlutils.CreateGlobalNotification(req)
}
