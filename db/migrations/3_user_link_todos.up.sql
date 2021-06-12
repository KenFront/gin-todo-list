ALTER TABLE todos ADD COLUMN user_id uuid;

ALTER TABLE todos ADD FOREIGN KEY (user_id) REFERENCES users(ID);
