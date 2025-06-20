package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rk5634/payment-system/internal/config"
	"github.com/rk5634/payment-system/internal/util"
)

func main() {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	config.LoadConfig(env)
	util.InitLogger(config.AppConfig.Log.Level)

	util.Logger.Info("Starting Payment System")
	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set Gin to release mode for production-like logs
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Middleware: recovery (panic handler), logging
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// Routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok",
			"message": "Hi dev - Payment service is running"})
	})

	// Start server
	log.Printf("🚀 Starting payment service on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
