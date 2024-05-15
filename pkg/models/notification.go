package models

import (
	"encoding/json"
	"time"
)

type GroupInviteExtra struct {
	GroupID     string `json:"groupId" validate:"required"`
	OrganizerID string `json:"organizerId" validate:"required"`
}
type GroupSettingsChangedExtra struct {
	GroupID string `json:"groupId" validate:"required"`
}
type FriendRequestExtra struct {
	UserID string `json:"userId" validate:"required"`
}
type EventInviteExtra struct {
	ActivityID  string `json:"activityId" validate:"required"`
	OrganizerID string `json:"organizerId" validate:"required"`
}
type EventApproachingExtra struct {
	ActivityID string `json:"activityId" validate:"required"`
	Date       string `json:"date" validate:"required"`
}
type CreateUserNotificationStruct struct {
	RecipientID string          `json:"recipientId" validate:"required"`
	Type        string          `json:"type" validate:"required"`
	Extra       json.RawMessage `json:"extra,omitempty"`
}
type Notification struct {
	ID          uint64    `json:"-"`
	IDString    string    `json:"id"`
	Expires     time.Time `json:"expires"`
	Extra       string    `json:"extra"`
	IsRead      bool      `json:"isRead"`
	RecipientID string    `json:"recipientId,omitempty"`
	Type        string    `json:"type"`
}
