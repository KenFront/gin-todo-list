CREATE TABLE IF NOT EXISTS todos (
    ID uuid NOT NULL,
    title varchar(100) NOT NULL,
    description varchar(65535) NOT NULL,
    status varchar(16) NOT NULL CHECK(status IN ('idle', 'completed')),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (ID)
);
