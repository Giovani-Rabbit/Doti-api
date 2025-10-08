-- name: CreateTask :one
INSERT INTO tasks (module_id, name, is_completed, position, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, module_id, name, is_completed, position, created_at, updated_at;

-- name: CheckTaskExists :one
SELECT EXISTS (
    SELECT 1 
    FROM tasks
    WHERE id = $1
) AS EXISTS;

-- name: FindTaskById :one
SELECT id, module_id, name, is_completed, position, created_at, updated_at
FROM tasks WHERE id = $1;

-- name: ListTasksByModuleId :many
SELECT id, module_id, name, is_completed, position, created_at, updated_at
FROM tasks WHERE module_id = $1
ORDER BY position ASC;

-- name: TaskPositionExists :one
SELECT EXISTS (
    SELECT 1 FROM tasks
    WHERE module_id = $1 
    AND position = $2 
) AS exists;

-- name: UpdateTaskCompletion :exec
UPDATE tasks
SET is_completed = $1
WHERE id = $2;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: FindOwnerIdByTaskId :one
SELECT m.user_id
FROM tasks t
INNER JOIN modules m ON m.id = t.module_id
WHERE t.id = $1;

-- name: UpdateTaskName :exec
UPDATE tasks
SET name = $2
WHERE id = $1;