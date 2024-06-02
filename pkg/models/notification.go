package models

import (
	"encoding/json"
	"time"
)

type NotificationDispatchReq struct {
	TargetRecipients []string               `json:"targetRecipients" validate:"required"`
	Extra            map[string]interface{} `json:"extra"`
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
