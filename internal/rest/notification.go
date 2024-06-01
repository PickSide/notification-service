package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notification-service/internal/service"
	"notification-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func DispatchFriendRequestNotification(g *gin.Context) {
	var req models.FriendRequestNotificationDispatchReq

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

}

func DispatchGroupInviteNotification(g *gin.Context) {
	var req models.GroupInviteNotificationDispatchReq

	if err := g.ShouldBindJSON(&req); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}
	log.Println("TargetRecipients", req.TargetRecipients)
	for _, targetID := range req.TargetRecipients {
		extraData, err := json.Marshal(map[string]string{
			"groupId":     req.GroupID,
			"organizerId": req.OrganizerID,
		})
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error marshaling extra data",
				"success": false,
			})
			continue
		}

		err = service.CreateNotification(models.CreateNotificationStruct{
			RecipientID: targetID,
			Type:        "group-invite",
			Extra:       json.RawMessage(extraData),
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
	var req models.GroupSettingsChangeNotificationDispatchReq

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

	result, err := service.GetUserNotifications(userKey)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"results": result,
		"success": true,
	})
}
