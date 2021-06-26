ALTER TABLE users ADD COLUMN deleted_at timestamp;
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_account_key;
CREATE UNIQUE INDEX users_account_deleted_at_key ON users(account) WHERE deleted_at IS NULL;