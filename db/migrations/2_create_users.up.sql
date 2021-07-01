CREATE TABLE IF NOT EXISTS users (
    ID uuid NOT NULL,
    name varchar(100) NOT NULL,
    account varchar(100) NOT NULL UNIQUE,
    password varchar(100) NOT NULL,
    email varchar(255) NOT NULL,
    status varchar(16) NOT NULL CHECK(status IN ('active', 'inactive', 'forbidden')),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (ID)
);