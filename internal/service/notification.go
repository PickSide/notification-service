package service

import (
	"encoding/json"
	"errors"
	sqlutils "notification-service/internal/sql"
	"notification-service/pkg/models"
	"strings"
)

func CreateUserNotification(req models.CreateUserNotificationStruct) (*string, error) {
	decoder := json.NewDecoder(strings.NewReader(string(req.Extra)))
	decoder.DisallowUnknownFields()

	switch req.Type {
	case "group-invite":
		var extra models.GroupInviteExtra
		err := decoder.Decode(&extra)
		if err != nil {
			return nil, errors.New("wrong payload for group invite. Missing GroupID or OrganizerID")
		}
	case "friend-request":
		var extra models.FriendRequestExtra
		err := decoder.Decode(&extra)
		if err != nil {
			return nil, errors.New("wrong payload for friend request. Missing UserID")
		}
	case "group-settings-changed":
		var extra models.GroupSettingsChangedExtra
		err := decoder.Decode(&extra)
		if err != nil {
			return nil, errors.New("wrong payload for group settings changes. Missing GroupID")
		}
	case "event-invite":
		var extra models.GroupSettingsChangedExtra
		err := decoder.Decode(&extra)
		if err != nil {
			return nil, errors.New("wrong payload for event invite. Missing ActivityID or organizerID")
		}
	case "event-approaching":
		var extra models.EventApproachingExtra
		err := decoder.Decode(&extra)
		if err != nil {
			return nil, errors.New("wrong payload for event approaching. Missing ActivityID or Date")
		}

	default:
		return nil, errors.New("unsupported notification type")
	}
	return sqlutils.CreateNotification(req)
}
func GetGlobalNotifications() (*[]models.GlobalNotification, error) {
	return sqlutils.GetGlobalNotifications()
}
func GetUserNotifications(userID string) (*[]models.Notification, error) {
	return sqlutils.GetUserNotifications(userID)
}
func UpdateSeenStatus(notificationID string) error {
	return sqlutils.UpdateSeenStatus(notificationID)
}
