package models

import (
	"encoding/json"
	"time"
)

type GroupInviteNotificationDispatchReq struct {
	GroupID          string   `json:"groupId" validate:"required"`
	OrganizerID      string   `json:"organizerId" validate:"required"`
	TargetRecipients []string `json:"targetRecipients" validate:"required"`
}
type GroupSettingsChangeNotificationDispatchReq struct {
	GroupID          string   `json:"groupId" validate:"required"`
	OrganizerID      string   `json:"organizerId" validate:"required"`
	SettingAffected  []string `json:"settingAffected" validate:"required"`
	TargetRecipients []string `json:"targetRecipients" validate:"required"`
}
type FriendRequestNotificationDispatchReq struct {
	GroupID          string   `json:"groupId" validate:"required"`
	OrganizerID      string   `json:"organizerId" validate:"required"`
	TargetRecipients []string `json:"targetRecipients" validate:"required"`
}
type CreateNotificationStruct struct {
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
