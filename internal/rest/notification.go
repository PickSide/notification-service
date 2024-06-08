package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notification-service/internal/service"
	"notification-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func DispatchFriendRequestNotification(g *gin.Context) {
	var req models.NotificationDispatchReq
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

}

func DispatchGroupInviteNotification(g *gin.Context) {
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

		err = service.CreateNotification(models.CreateNotificationStruct{
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

func DispatchGroupSettingsChangeNotification(g *gin.Context) {
	var req models.NotificationDispatchReq

	if err := g.ShouldBindJSON(&req); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

}

func GetNotifications(g *gin.Context) {
	userKey := g.Query("userKey")

	if userKey == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
	}

	results, err := service.GetUserNotifications(userKey)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, results)
}

func DeleteNotification(g *gin.Context) {
	notificationID := g.Param("notificationID")

	if notificationID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
	}

	err := service.DeleteNotification(notificationID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, nil)
}
