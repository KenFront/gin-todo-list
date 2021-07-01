ALTER TABLE
    todos
ADD
    COLUMN user_id uuid,
ADD
    FOREIGN KEY (user_id) REFERENCES users(ID);