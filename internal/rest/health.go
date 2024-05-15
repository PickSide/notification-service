package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"service": "notification-service",
		"status":  "up",
	})
}
