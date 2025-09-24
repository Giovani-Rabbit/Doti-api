-- name: CreateTask :one
INSERT INTO tasks (module_id, name, is_completed, position, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, module_id, name, is_completed, position, created_at, updated_at;

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

-- name: SwapTaskPosition :exec
UPDATE tasks 
SET position = CASE id
    WHEN @task_id_1 THEN @position_1::int
    WHEN @task_id_2 THEN @position_2::int
END
WHERE id IN (@task_id_1, @task_id_2);