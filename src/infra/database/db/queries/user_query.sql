-- name: CreateUser :exec
INSERT INTO users (ID, Email, Name, Password)
VALUES ($1, $2, $3, $4);