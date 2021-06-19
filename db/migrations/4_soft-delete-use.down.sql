ALTER TABLE users 
    DROP CONSTRAINT IF EXISTS users_account_deleted_at_key,
    ADD UNIQUE (account);
ALTER TABLE users DROP COLUMN IF EXISTS deleted_at;
