package api

import (
	"net/http"
	"notification-service/internal/database"

	"github.com/gin-gonic/gin"
)

func GetNotifications(g *gin.Context) {
	userKey := g.Query("userKey")

	if userKey == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
	}

	results, err := database.GetUserNotifications(userKey)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, results)
}
