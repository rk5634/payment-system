package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	httpHandler "github.com/rk5634/payment-system/internal/handler/http"
	"github.com/rk5634/payment-system/internal/service"
)

// RegisterRoutes attaches all HTTP routes to the given Gin engine.
func RegisterRoutes(router *gin.Engine, svc *service.Service) {
	log.Println("Registering API routes...")
	v1 := router.Group("/api/v1")
	log.Println("Grouped routes under /api/v1")

	// Register health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Hi dev - Payment service is running",
		})
	})

	// Register payment routes
	paymentHandler := httpHandler.NewPaymentHandler(svc)
	paymentHandler.RegisterRoutes(v1)
}
