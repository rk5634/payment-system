package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rk5634/payment-system/internal/domain/payment"

	"github.com/google/uuid"
)

// Service defines business logic for managing payments.
type Service struct {
	repo payment.Repository
}

// NewService returns a new PaymentService.
func NewService(repo payment.Repository) *Service {
	return &Service{repo: repo}
}

// CreatePayment creates a new payment using the input request.
func (s *Service) CreatePayment(ctx context.Context, req *payment.CreatePaymentRequest) (*payment.Payment, error) {
	now := time.Now().UTC()

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id: %w", err)
	}

	p := &payment.Payment{
		// ID is omitted, Postgres will generate and return it
		UserID:         userUUID,
		Amount:         req.Amount,
		Currency:       req.Currency,
		Status:         "pending",
		PaymentMethod:  req.PaymentMethod,
		TransactionID:  nil,              // optional, will remain NULL unless set
		Description:    req.Description,  // optional
		Metadata:       req.Metadata,     // json.RawMessage or nil
		CompletedAt:    nil,              // will remain NULL until marked completed
		FailedAt:       nil,
		RefundedAt:     nil,
		IsDeleted:      false,
		Version:        1,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// Insert into DB and populate p.ID from RETURNING clause
	if err := s.repo.CreatePayment(ctx, p); err != nil {
		return nil, fmt.Errorf("create payment failed: %w", err)
	}

	return p, nil
}

