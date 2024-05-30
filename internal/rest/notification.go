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
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

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
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

}

func GetUserNotifications(g *gin.Context) {
	userID := g.Param("userID")

	result, err := service.GetUserNotifications(userID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"results": result,
		"success": true,
	})
}

// func CreateUserNotification(g *gin.Context) {
// 	var req models.CreateUserNotificationStruct

// 	if err := g.ShouldBindJSON(&req); err != nil {
// 		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err := service.CreateUserNotification(req)
// 	if err != nil {
// 		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

//		g.JSON(http.StatusOK, gin.H{
//			"message": "Notification successfully created",
//			"success": true,
//		})
//	}

// func UpdateSeenStatus(g *gin.Context) {
// 	notificationID := g.Param("notificationID")

// 	err := service.UpdateSeenStatus(notificationID)
// 	if err != nil {
// 		g.JSON(http.StatusNotFound, gin.H{
// 			"error":   err.Error(),
// 			"message": "UpdateSeenStatus - Failed to update seen status of notification",
// 			"success": false,
// 		})
// 		return
// 	}
// 	g.JSON(http.StatusNoContent, gin.H{
// 		"success": true,
// 	})
// }
