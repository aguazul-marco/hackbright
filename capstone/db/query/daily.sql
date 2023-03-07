-- name: CreateDailyGoal :one
INSERT INTO daily_goals (
  discription,
  user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUserDailyGoal :one
SELECT * FROM daily_goals
WHERE user_id = $1 LIMIT 1;

-- name: DeleteDailyGoal :exec
DELETE FROM daily_goals
WHERE id = $1;

-- name: DailyCompleteStatusUpdate :one
UPDATE daily_goals
SET completed = $2
WHERE id = $1
RETURNING *;

-- name: DailyUserGoals :many
SELECT * FROM daily_goals
WHERE user_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: DailyCompletedGoals :many
SELECT * FROM daily_goals
WHERE user_id = $1 AND completed = true
ORDER BY created_at;

-- name: DailyUncompletedGoals :many
SELECT * FROM daily_goals
WHERE user_id = $1 AND completed = false
ORDER BY created_at;