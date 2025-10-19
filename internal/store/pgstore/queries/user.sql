-- name: CreateUser :one
INSERT INTO users ("name", "email", "password_hash")
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetUserById :one
SELECT id, name, password_hash, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, name, password_hash, created_at, updated_at
FROM users
WHERE email = $1;