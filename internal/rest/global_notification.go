package rest

import (
	"net/http"
	"notification-service/internal/service"
	"notification-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateGlobalNotification(g *gin.Context) {
	var req models.CreateGlobalNotificationRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err := service.CreateGlobalNotification(req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Notification successfully created",
		"success": true,
	})
}
func GetGlobalNotifications(g *gin.Context) {
	result, err := service.GetGlobalNotifications()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"results": result})
}
