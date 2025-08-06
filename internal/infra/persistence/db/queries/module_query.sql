-- name: CreateModule :one
INSERT INTO modules (id, user_id, Name, is_open, icon, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, name, is_open, icon, created_at, updated_at;

-- name: ListModuleByUserID :many
SELECT id, user_id, name, is_open, icon, created_at, updated_at 
FROM modules
WHERE user_id = $1;

-- name: UpdateModuleName :exec
UPDATE modules 
SET name = $2
WHERE id = $1;