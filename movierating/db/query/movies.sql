-- name: CreateMovie :one
INSERT INTO movies (
  title, 
  overview,
  release_date,
  poster_url
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetMovie :one
SELECT * FROM movies
WHERE id = $1 LIMIT 1;

-- name: GetMovies :many
SELECT * FROM movies
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteMovie :exec
DELETE FROM movies
WHERE id = $1;