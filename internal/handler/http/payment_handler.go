package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk5634/payment-system/internal/domain/payment"
	"github.com/rk5634/payment-system/internal/service"
)

// PaymentHandler holds dependencies for HTTP routes.
type PaymentHandler struct {
	svc *service.Service
}

// NewPaymentHandler creates a new PaymentHandler.
func NewPaymentHandler(svc *service.Service) *PaymentHandler {
	return &PaymentHandler{svc: svc}
}

// RegisterRoutes registers payment routes to the given Gin router group.
func (h *PaymentHandler) RegisterRoutes(rg *gin.RouterGroup) {
	log.Println("Registering payment routes...")
	rg.POST("/payments", h.CreatePayment)
}

// CreatePayment handles POST /payments requests.
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req payment.CreatePaymentRequest

	// Bind JSON body into struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
		return
	}

	// Call service to create payment
	p, err := h.svc.CreatePayment(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create payment: " + err.Error()})
		return
	}

	// Return created payment
	c.JSON(http.StatusCreated, p)
}
