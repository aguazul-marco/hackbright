-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2, last_name = $3
WHERE id = $1
RETURNING *;