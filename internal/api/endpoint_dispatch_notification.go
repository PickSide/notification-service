package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notification-service/internal/database"
	"notification-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func DispatchNotification(g *gin.Context) {
	var req models.NotificationDispatchReq

	if err := g.ShouldBindJSON(&req); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

	for _, targetID := range req.TargetRecipients {
		jsonData, err := json.Marshal(req.Extra)
		if err != nil {
			jsonData = []byte{}
		}

		err = database.CreateNotification(models.CreateNotificationStruct{
			RecipientID: targetID,
			Type:        "group-invite",
			Extra:       jsonData,
		})
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Errorf("error dispatching notification for user: %s", targetID),
				"success": false,
			})
			continue
		}
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Group invites disaptched successfully",
		"success": true,
	})
}
