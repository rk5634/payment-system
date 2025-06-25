-- ======================================
-- PAYMENTS TABLE SCHEMA (Production-Ready)
-- ======================================

CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique payment identifier
    user_id UUID NOT NULL,                                -- Reference to the user (add FK as needed)
    amount NUMERIC(10, 2) NOT NULL,                       -- Payment amount (up to 99,999,999.99)
    currency CHAR(3) NOT NULL,                            -- ISO 4217 currency code (e.g., USD, EUR)
    status VARCHAR(50) NOT NULL DEFAULT 'pending',        -- Payment status (pending, completed, failed, refunded)
    payment_method VARCHAR(50) NOT NULL,                  -- Payment method (e.g., card, PayPal)
    transaction_id UUID UNIQUE,                           -- External transaction reference (nullable, unique)
    description TEXT,                                     -- Optional notes (e.g., invoice number)
    metadata JSONB,                                       -- Flexible key-value data for extra info
    completed_at TIMESTAMPTZ,                             -- When payment completed (nullable)
    failed_at TIMESTAMPTZ,                                -- When payment failed (nullable)
    refunded_at TIMESTAMPTZ,                              -- When refunded (nullable)
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,            -- Soft delete flag
    version INT NOT NULL DEFAULT 1,                       -- Version for optimistic locking (optional)
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),        -- Record creation timestamp (UTC)
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()         -- Record last update timestamp (UTC)
);

-- ======================================
-- INDEXES
-- ======================================

-- Speed up queries by user
CREATE INDEX IF NOT EXISTS idx_payments_user_id
ON payments (user_id);

-- Speed up queries by status
CREATE INDEX IF NOT EXISTS idx_payments_status
ON payments (status);

-- Speed up soft delete lookups
CREATE INDEX IF NOT EXISTS idx_payments_is_deleted
ON payments (is_deleted);

-- Speed up created_at queries (e.g. for reports / partitioning)
CREATE INDEX IF NOT EXISTS idx_payments_created_at
ON payments (created_at);

-- ======================================
-- TRIGGER: auto-update updated_at on change
-- ======================================

CREATE OR REPLACE FUNCTION payments_set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_payments_set_updated_at
BEFORE UPDATE ON payments
FOR EACH ROW
EXECUTE FUNCTION payments_set_updated_at();

-- ======================================
-- (OPTIONAL) FOREIGN KEY (uncomment if users table exists)
-- ======================================
-- ALTER TABLE payments
-- ADD CONSTRAINT fk_payments_user
-- FOREIGN KEY (user_id) REFERENCES users(id)
-- ON DELETE CASCADE;

-- ======================================
-- (OPTIONAL) ENUM constraints (replace VARCHAR with ENUM if preferred)
-- ======================================
-- You could create ENUM types:
-- CREATE TYPE payment_status AS ENUM ('pending', 'completed', 'failed', 'refunded');
-- CREATE TYPE payment_method AS ENUM ('card', 'paypal', 'upi', 'bank_transfer');
--
-- And use in table:
-- status payment_status NOT NULL DEFAULT 'pending',
-- payment_method payment_method NOT NULL,
