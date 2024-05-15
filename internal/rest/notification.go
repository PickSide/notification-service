package rest

import (
	"net/http"
	"notification-service/internal/service"
	"notification-service/pkg/models"

	"github.com/gin-gonic/gin"
)

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
func CreateUserNotification(g *gin.Context) {
	var req models.CreateUserNotificationStruct

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err := service.CreateUserNotification(req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Notification successfully created",
		"success": true,
	})
}
