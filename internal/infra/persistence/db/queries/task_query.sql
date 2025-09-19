-- name: CreateTask :one
INSERT INTO tasks (module_id, name, is_completed, position, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, module_id, name, is_completed, position, created_at, updated_at;

-- name: ListTasksByModuleId :many
SELECT id, module_id, name, is_completed, position, created_at, updated_at
FROM tasks WHERE module_id = $1
ORDER BY position ASC;

-- name: GetTaskbyId :one
SELECT id, module_id, name, is_completed, position, created_at, updated_at
FROM tasks WHERE id = $1;

-- name: GetTaskByPosition :one
SELECT id, module_id, name, is_completed, position, created_at, updated_at
FROM tasks WHERE module_id = $1 AND position = $2;

-- name: UpdateTaskPosition :exec
UPDATE tasks SET position = $2
WHERE id = $1;