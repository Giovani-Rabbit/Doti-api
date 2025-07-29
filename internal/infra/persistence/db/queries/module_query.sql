-- name: CreateModule :one
INSERT INTO modules (id, user_id, Name, is_open, icon, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, name, is_open, icon;