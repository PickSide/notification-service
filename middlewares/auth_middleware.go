package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyBearer() gin.HandlerFunc {
	return func(g *gin.Context) {
		token := strings.Split(g.GetHeader("Authorization"), " ")[1]

		req, err := http.NewRequest("GET", os.Getenv("AUTHENTICATION_SERVER_BASE_URL")+"/verify-token", nil)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Failed to connect to auth-service",
				"success": false,
			})
			return
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Can't call the auth-service",
				"success": false,
			})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is invalid or expired",
				"success": false,
			})
			return
		}

		g.Next()
	}
}
