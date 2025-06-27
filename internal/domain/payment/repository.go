package payment

import "context"

// Repository defines the interface for payment-related persistence operations.
type Repository interface {
    CreatePayment(ctx context.Context, p *Payment) error
    GetPaymentByID(ctx context.Context, id string) (*Payment, error)
    UpdatePaymentStatus(ctx context.Context, id string, status string) error
    DeletePayment(ctx context.Context, id string) error

    // Optional advanced queries
    ListPaymentsByUser(ctx context.Context, userID string) ([]*Payment, error)
    SoftDeletePayment(ctx context.Context, id string) error
}
