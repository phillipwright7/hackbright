-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    email
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUserDetails :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;