package database

import (
	"errors"
	"notification-service/pkg/models"
	"notification-service/pkg/utils"
	"time"
)

const (
	DAYS_TO_EXPIRE = 4
)

func CreateNotification(req models.CreateNotificationStruct) error {
	uint64RecipientID, err := utils.StringToUint64(req.RecipientID)
	if err != nil {
		return errors.New("CreateNotification - Error during conversion (Malformed ID)")
	}

	tx, err := GetClient().Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	_, err = tx.Exec(`INSERT INTO notifications (expires, extra, recipient_id, type) VALUES (?, ?, ?, ?)`,
		time.Now().AddDate(0, 0, DAYS_TO_EXPIRE), req.Extra, uint64RecipientID, req.Type,
	)

	return err
}
