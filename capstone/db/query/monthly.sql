-- name: CreateMonthlyGoal :one
INSERT INTO monthly_goals (
  discription,
  user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUserMonthlyGoal :one
SELECT * FROM monthly_goals
WHERE user_id = $1 LIMIT 1;

-- name: DeleteMonthlyGoal :exec
DELETE FROM monthly_goals
WHERE id = $1;

-- name: MonthlyCompleteStatusUpdate :one
UPDATE monthly_goals
SET completed = $2
WHERE id = $1
RETURNING *;

-- name: MonthlyUserGoals :many
SELECT * FROM monthly_goals
WHERE user_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: UserMonthlyCompletedGoals :many
SELECT * FROM monthly_goals
WHERE user_id = $1 AND completed = true
ORDER BY created_at;

-- name: UserMonthlyUncompletedGoals :many
SELECT * FROM monthly_goals
WHERE user_id = $1 AND completed = false
ORDER BY created_at;