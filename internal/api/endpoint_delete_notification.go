package api

import (
	"net/http"
	"notification-service/internal/database"

	"github.com/gin-gonic/gin"
)

func DeleteNotification(g *gin.Context) {
	notificationID := g.Param("notificationID")

	if notificationID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
	}

	err := database.DeleteNotification(notificationID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, nil)
}
