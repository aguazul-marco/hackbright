-- name: CreateWeeklyGoal :one
INSERT INTO weekly_goals (
  discription,
  user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUserWeeklyGoal :one
SELECT * FROM weekly_goals
WHERE user_id = $1 LIMIT 1;

-- name: DeleteWeeklyGoal :exec
DELETE FROM weekly_goals
WHERE id = $1;

-- name: WeeklyCompleteStatusUpdate :one
UPDATE weekly_goals
SET completed = $2
WHERE id = $1
RETURNING *;

-- name: WeeklyUserGoals :many
SELECT * FROM weekly_goals
WHERE user_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: WeeklyCompletedGoals :many
SELECT * FROM weekly_goals
WHERE user_id = $1 AND completed = true
ORDER BY created_at;

-- name: WeeklyUncompletedGoals :many
SELECT * FROM weekly_goals
WHERE user_id = $1 AND completed = false
ORDER BY created_at;