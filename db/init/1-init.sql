CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE todo_status AS ENUM ('idle', 'completed');

CREATE TABLE IF NOT EXISTS todos (
    id uuid DEFAULT uuid_generate_v4 (),
    title varchar(100) NOT NULL,
    description varchar(65535) NOT NULL,
    status todo_status NOT NULL DEFAULT 'idle'
);