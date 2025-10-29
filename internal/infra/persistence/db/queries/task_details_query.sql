-- name: CreateTaskDetails :exec
INSERT INTO task_details (task_id, description)
VALUES ($1, $2);

-- name: UpdateTaskDetailsDescription :execrows
UPDATE task_details
SET description = $2
WHERE task_id = $1;

-- name: UpdateTaskDetailsPomodoroTarget :execrows
UPDATE task_details
SET pomodoro_target = $2
WHERE task_id = $1;