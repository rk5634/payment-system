package payment

import (
    "time"

    "github.com/google/uuid"
    "encoding/json"
)

// Payment represents the domain model for a payment transaction.
type Payment struct {
    ID             uuid.UUID       `json:"id" db:"id"`
    UserID         uuid.UUID       `json:"user_id" db:"user_id"`
    Amount         float64         `json:"amount" db:"amount"`               // numeric(10,2)
    Currency       string          `json:"currency" db:"currency"`           // ISO 4217
    Status         string          `json:"status" db:"status"`               // pending, completed, etc.
    PaymentMethod  string          `json:"payment_method" db:"payment_method"`
    TransactionID  *uuid.UUID      `json:"transaction_id,omitempty" db:"transaction_id"` // optional
    Description    *string         `json:"description,omitempty" db:"description"`       // optional
    Metadata       json.RawMessage `json:"metadata,omitempty" db:"metadata"`             // flexible key-value JSON
    CompletedAt    *time.Time      `json:"completed_at,omitempty" db:"completed_at"`     // nullable
    FailedAt       *time.Time      `json:"failed_at,omitempty" db:"failed_at"`
    RefundedAt     *time.Time      `json:"refunded_at,omitempty" db:"refunded_at"`
    IsDeleted      bool            `json:"is_deleted" db:"is_deleted"`
    Version        int             `json:"version" db:"version"`
    CreatedAt      time.Time       `json:"created_at" db:"created_at"`
    UpdatedAt      time.Time       `json:"updated_at" db:"updated_at"`
}


type CreatePaymentRequest struct {
	UserID         string       `json:"user_id" db:"user_id"`
    Amount         float64         `json:"amount" db:"amount"`               // numeric(10,2)
    Currency       string          `json:"currency" db:"currency"`           // ISO 4217
    Status         string          `json:"status" db:"status"`               // pending, completed, etc.
    PaymentMethod  string          `json:"payment_method" db:"payment_method"`
    TransactionID  *uuid.UUID      `json:"transaction_id,omitempty" db:"transaction_id"` // optional
    Description    *string         `json:"description,omitempty" db:"description"`       // optional
    Metadata       json.RawMessage `json:"metadata,omitempty" db:"metadata"`             // flexible key-value JSON
    CompletedAt    *time.Time      `json:"completed_at,omitempty" db:"completed_at"`     // nullable
    FailedAt       *time.Time      `json:"failed_at,omitempty" db:"failed_at"`
    RefundedAt     *time.Time      `json:"refunded_at,omitempty" db:"refunded_at"`
    IsDeleted      bool            `json:"is_deleted" db:"is_deleted"`
    Version        int             `json:"version" db:"version"`
    CreatedAt      time.Time       `json:"created_at" db:"created_at"`
    UpdatedAt      time.Time       `json:"updated_at" db:"updated_at"`
}
