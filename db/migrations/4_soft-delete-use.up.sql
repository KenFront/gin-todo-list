ALTER TABLE users ADD COLUMN deleted_at timestamp;
ALTER TABLE users
    DROP CONSTRAINT IF EXISTS users_account_key,
    ADD UNIQUE (account, deleted_at);
