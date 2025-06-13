-- name: CreateUser :one
INSERT INTO users (ID, Email, Name, Password, Created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING ID, Email, Name, Created_at;

-- name: CheckUserExists :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE Email = $1
) AS EXISTS;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE Email = $1 LIMIT 1;

-- name: FindUserByEmailAndPassword :one
SELECT * FROM users WHERE Email = $1 AND Password = $2 LIMIT 1;
