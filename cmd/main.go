package main

import (
	"log"
	"notification-service/internal/api"
	"notification-service/internal/database"
	"notification-service/internal/vault"
	"notification-service/pkg/env"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Env intialization
	env.InitalizeEnvs()

	// Vault initialization
	vault.InitializeVault()

	// Database initialization
	database.InitializeDB(vault.Envars["DSN"].(string))
	defer database.Close()

	gin.SetMode(env.GIN_MODE)
	g := gin.Default()

	g.Use(cors.New(buildCors()))

	g.GET("/health", api.GetHealth)

	g.GET("/notifications", api.GetNotifications)
	g.POST("/notifications/dispatch", api.DispatchNotification)
	g.PUT("/notifications/seen", api.DispatchNotification)
	g.DELETE("/notifications/:notificationID", api.DeleteNotification)

	PrintServiceInformation()

	if err := g.Run(":8084"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func PrintServiceInformation() {
	log.Printf("Mode %s", env.GIN_MODE)
	log.Printf("Service name: %s", env.SERVICE_NAME)
	log.Printf("Version: %s", env.VERSION)
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
