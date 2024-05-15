package models

import "time"

type CreateGlobalNotificationRequest struct {
	Content string `json:"content"`
	Title   string `json:"title,omitempty"`
}

type GlobalNotification struct {
	ID        uint64    `json:"-"`
	IDString  string    `json:"id"`
	Content   string    `json:"content"`
	Expires   time.Time `json:"expires"`
	Title     string    `json:"title,omitempty"`
	CreatedAt string    `json:"-"`
}
