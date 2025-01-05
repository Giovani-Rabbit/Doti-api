-- name: CreateUser :exec
INSERT INTO users (ID, Email, Name, Password)
VALUES ($1, $2, $3, $4);

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
