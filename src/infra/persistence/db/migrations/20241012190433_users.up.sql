CREATE TABLE IF NOT EXISTS users (
    id UUID    NOT NULL PRIMARY KEY,
    email      VARCHAR(255)             NOT NULL,
    name       VARCHAR(255)             NOT NULL,
    password   VARCHAR(255)             NOT NULL,
    is_admin   BOOLEAN                  NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
)