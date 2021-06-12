CREATE TYPE todo_status AS ENUM ('idle', 'completed');

CREATE TABLE IF NOT EXISTS todos (
    ID uuid NOT NULL,
    title varchar(100) NOT NULL,
    description varchar(65535) NOT NULL,
    status todo_status NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (ID)
);
