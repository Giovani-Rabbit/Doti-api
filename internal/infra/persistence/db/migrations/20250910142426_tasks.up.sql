CREATE TABLE IF NOT EXISTS tasks (
    id              INTEGER GENERATED ALWAYS AS IDENTITY    NOT NULL    PRIMARY KEY,
    module_id       INTEGER                                 NOT NULL,
    name            VARCHAR(255)                            NOT NULL,
    is_completed    BOOLEAN                                 NOT NULL    DEFAULT false,
    position        INTEGER                                 NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE                NOT NULL    DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE                NOT NULL    DEFAULT now(),

    CONSTRAINT fk_module FOREIGN KEY (module_id) REFERENCES modules(id)
)