CREATE TABLE IF NOT EXISTS task_details (
    id              INTEGER GENERATED ALWAYS AS IDENTITY    NOT NULL    PRIMARY KEY,
    task_id         INTEGER                                 NOT NULL,
    description     TEXT,
    pomodoro_target INTEGER                                             DEFAULT 0,
    pomodoro_spent  INTEGER                                             DEFAULT 0,
    total_time      INTEGER                                             DEFAULT 0,
    updated_at      TIMESTAMP WITH TIME ZONE                NOT NULL    DEFAULT now(),

    CONSTRAINT fk_task FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
)