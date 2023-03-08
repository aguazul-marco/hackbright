-- name: CreateMonthlyGoal :one
INSERT INTO monthly_goals (
  discription,
  user_id,
  completed
) VALUES (
  $1, $2, false
) RETURNING *;

-- name: DeleteMonthlyGoal :exec
DELETE FROM monthly_goals
WHERE id = $1;

-- name: MonthlyCompleteStatusUpdate :one
UPDATE monthly_goals
SET completed = $2
WHERE id = $1
RETURNING *;

-- name: UserMonthlyGoals :many
SELECT * FROM monthly_goals
WHERE user_id = $1
ORDER BY created_at;

-- name: UserMonthlyCompletedGoals :many
SELECT * FROM monthly_goals
WHERE user_id = $1 AND completed = true
ORDER BY created_at;

-- name: UserMonthlyUncompletedGoals :many
SELECT * FROM monthly_goals
WHERE user_id = $1 AND completed = false
ORDER BY created_at;