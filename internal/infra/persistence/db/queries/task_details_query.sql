-- name: CreateTaskDetails :exec
INSERT INTO task_details (task_id, description)
VALUES ($1, $2);