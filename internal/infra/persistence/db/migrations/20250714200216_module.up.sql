CREATE TABLE IF NOT EXISTS modules (
    id         SERIAL                   NOT NULL    PRIMARY KEY,
    user_id    UUID                     NOT NULL,
    name       VARCHAR(255)             NOT NULL,
    is_open    BOOLEAN                  NOT NULL    DEFAULT false,
    icon       VARCHAR(255)             NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL    DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL    DEFAULT now(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
)