package main

import (
	"log"
	"notification-service/internal/database"
	"notification-service/internal/rest"
	"notification-service/internal/vault"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := vault.Envars["DSN"].(string)
	database.Initialize(dsn)
	defer database.Close()

	gin.SetMode(os.Getenv("GIN_MODE"))
	g := gin.Default()

	g.Use(cors.New(buildCors()))

	g.GET("/health", rest.GetHealth)

	g.GET("/notifications", rest.GetNotifications)
	g.DELETE("/notifications/:notificationID", rest.DeleteNotification)

	g.POST("/dispatch/friend-request", rest.DispatchFriendRequestNotification)
	g.POST("/dispatch/group-invite", rest.DispatchGroupInviteNotification)
	g.POST("/dispatch/group-settings-change", rest.DispatchGroupSettingsChangeNotification)

	PrintServiceInformation()

	if err := g.Run(":8084"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func PrintServiceInformation() {
	log.Printf("Mode %s", os.Getenv("GIN_MODE"))
	log.Printf("Service name: %s", os.Getenv("SERVICE_NAME"))
	log.Printf("Version: %s", os.Getenv("SERVICE_VERSION"))
}

func buildCors() cors.Config {
	c := cors.DefaultConfig()
	c.AllowAllOrigins = false
	c.AllowCredentials = true
	c.AllowHeaders = []string{"Accept-Version", "Authorization", "Content-Type", "Origin", "X-Client-Version", "X-CSRF-Token", "X-Request-Id"}
	c.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	c.AllowWebSockets = true
	c.MaxAge = 24 * time.Hour

	c.AllowOriginFunc = func(origin string) bool {
		return true
	}
	return c
}
