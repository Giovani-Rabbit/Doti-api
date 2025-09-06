-- name: CreateModule :one
INSERT INTO modules (user_id, Name, is_open, icon, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, user_id, name, is_open, icon, created_at, updated_at;

-- name: ListModuleByUserID :many
SELECT id, user_id, name, is_open, icon, created_at, updated_at 
FROM modules
WHERE user_id = $1 ORDER BY created_at ASC;

-- name: UpdateModuleName :exec
UPDATE modules 
SET name = $2
WHERE id = $1;

-- name: DeleteModule :exec
DELETE FROM modules WHERE id = $1;

-- name: CheckModuleExists :one
SELECT EXISTS (
    SELECT 1 
    FROM modules 
    WHERE id = $1
) AS EXISTS;

-- name: UpdateIcon :exec
UPDATE modules 
SET icon = $1
WHERE id = $2;