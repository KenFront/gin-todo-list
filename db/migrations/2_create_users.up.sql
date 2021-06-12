CREATE TYPE user_status AS ENUM ('active', 'inactive', 'forbidden');

CREATE TABLE IF NOT EXISTS users (
    ID uuid NOT NULL,
    name varchar(100) NOT NULL,
    account varchar(100) NOT NULL UNIQUE,
    password varchar(100) NOT NULL,
    email varchar(255) NOT NULL,
    status user_status NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (ID)
);
