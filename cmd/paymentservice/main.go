package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/rk5634/payment-system/internal/api"
	"github.com/rk5634/payment-system/internal/config"
	dbpostgres "github.com/rk5634/payment-system/internal/database/postgres"
	repopostgres "github.com/rk5634/payment-system/internal/repository/postgres"
	"github.com/rk5634/payment-system/internal/service"
	"github.com/rk5634/payment-system/internal/util"
)

func main() {
	// Load environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Load config
	config.LoadConfig(env)
	util.InitLogger(config.AppConfig.Log.Level)

	// PostgreSQL configuration
	dbCfg := dbpostgres.Config{
		Host:     config.AppConfig.Database.PostgresHost,
		Port:     config.AppConfig.Database.PostgresPort,
		User:     config.AppConfig.Database.PostgresUser,
		Password: config.AppConfig.Database.PostgresPassword,
		DBName:   config.AppConfig.Database.PostgresDBName,
		SSLMode:  config.AppConfig.Database.PostgresSSLMode,
	}

	// Initialize DB connection
	log.Println("Connecting to PostgreSQL database...")
	db, err := dbpostgres.NewDB(dbCfg)
	if err != nil {
		log.Fatalf("❌ Failed to initialize PostgreSQL database: %v", err)
	}
	defer db.Close()
	log.Println("✅ Connected to PostgreSQL!")

	// Setup repository & service
	paymentRepo := repopostgres.NewPaymentRepository(db.Pool)
	paymentService := service.NewService(paymentRepo)

	// Setup Gin
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	// Register all routes (v1 API + health check)
	log.Println("Registering routes...")
	api.RegisterRoutes(router, paymentService)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Starting payment service on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
	log.Println("✅ Payment service is running!")
	log.Println("Visit http://localhost:" + port + "/health to check service status")
}
