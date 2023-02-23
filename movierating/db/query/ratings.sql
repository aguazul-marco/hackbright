-- name: CreateRating :one
INSERT INTO ratings (
  score,
  movie_id,
  user_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetRating :one
SELECT * FROM ratings
WHERE id = $1 LIMIT 1;

-- name: GetMovieRatings :many
SELECT * FROM ratings
WHERE movie_id = $1;

-- name: GetUserRatings :many
SELECT * FROM ratings
WHERE user_id = $1;

-- name: UpdateRating :one
UPDATE ratings
SET score = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRating :exec
DELETE FROM ratings
WHERE id = $1;