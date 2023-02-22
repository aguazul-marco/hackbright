-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    email
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;