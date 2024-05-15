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

	ns := g.Group("/notification-service")

	ns.GET("/health", rest.GetHealth)
	ns.GET("/global", rest.GetGlobalNotifications)
	ns.GET("/user/:userID", rest.GetUserNotifications)
	ns.POST("/user", rest.CreateUserNotification)
	ns.POST("/global", rest.CreateGlobalNotification)

	PrintServiceInformation()

	if err := g.Run(":8084"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func PrintServiceInformation() {
	log.Printf("Mode %s", os.Getenv("GIN_MODE"))
	log.Printf("Service name: %s", os.Getenv("GIN_MODE"))
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
