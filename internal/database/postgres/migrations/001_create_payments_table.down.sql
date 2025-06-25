-- ======================================
-- DOWN MIGRATION: Drop payments table and related structures
-- ======================================

-- Drop trigger first
DROP TRIGGER IF EXISTS trg_payments_set_updated_at ON payments;

-- Drop trigger function
DROP FUNCTION IF EXISTS payments_set_updated_at;

-- Drop indexes (not strictly necessary, they go with table, but explicit is clean)
DROP INDEX IF EXISTS idx_payments_user_id;
DROP INDEX IF EXISTS idx_payments_status;
DROP INDEX IF EXISTS idx_payments_is_deleted;
DROP INDEX IF EXISTS idx_payments_created_at;

-- Drop the payments table
DROP TABLE IF EXISTS payments;
