-- name: CreateUser :execresult
INSERT INTO user (
    id,
    name,
    email,
    password
)
VALUES (?,?,?,?);

-- name: GetUserById :one
SELECT * FROM user WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM user WHERE email = ?;