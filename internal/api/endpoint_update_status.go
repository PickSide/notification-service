package api

import (
	"net/http"
	"notification-service/internal/database"

	"github.com/gin-gonic/gin"
)

func SeenNotification(g *gin.Context) {
	notificationID := g.Query("notificationID")

	if notificationID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
	}

	err := database.UpdateSeenStatus(notificationID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusNoContent, nil)
}
