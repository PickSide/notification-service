package main

import (
	"log"
	"notification-service/internal/database"
	"notification-service/internal/vault"
	"notification-service/pkg/models"
	"notification-service/pkg/utils"
	"time"
)

func main() {
	dsn := vault.Envars["DSN"].(string)
	database.Initialize(dsn)
	defer database.Close()

	log.Println("Seeding table for notification-service...")

	notificationsData := []models.Notification{
		{RecipientID: "1", Type: "group-invite", Extra: `{"groupId": "1", "organizerId": "2"}`},
		{RecipientID: "1", Type: "group-settings-changed", Extra: `{"groupId": "1"}`},
		{RecipientID: "1", Type: "friend-request", Extra: `{"userId": "3"}`},
		{RecipientID: "1", Type: "event-invite", Extra: `{"activityId": "1", "organizerId": "2"}`},
		{RecipientID: "1", Type: "event-approaching", Extra: `{"activityId": "1", "date": "2024-04-18 04:35:01"}`},
	}

	for _, notification := range notificationsData {
		uint64RecipientID, err := utils.StringToUint64(notification.RecipientID)
		if err != nil {
			continue
		}
		_, err = database.GetClient().Exec(`INSERT INTO notifications (expires, extra, is_read, recipient_id, type) VALUES (?, ?, ?, ?, ?)`,
			time.Now().AddDate(0, 0, 1),
			notification.Extra,
			0,
			uint64RecipientID,
			notification.Type,
		)
		if err != nil {
			panic(err)
		}
	}

	log.Println("Done")
}
