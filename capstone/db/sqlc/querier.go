// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateDailyGoal(ctx context.Context, arg CreateDailyGoalParams) (DailyGoal, error)
	CreateMonthlyGoal(ctx context.Context, arg CreateMonthlyGoalParams) (MonthlyGoal, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateWeeklyGoal(ctx context.Context, arg CreateWeeklyGoalParams) (WeeklyGoal, error)
	DailyCompleteStatusUpdate(ctx context.Context, arg DailyCompleteStatusUpdateParams) (DailyGoal, error)
	DailyCompletedGoals(ctx context.Context, userID sql.NullInt32) ([]DailyGoal, error)
	DailyUncompletedGoals(ctx context.Context, userID sql.NullInt32) ([]DailyGoal, error)
	DeleteDailyGoal(ctx context.Context, id int64) error
	DeleteMonthlyGoal(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	DeleteWeeklyGoal(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (User, error)
	MonthlyCompleteStatusUpdate(ctx context.Context, arg MonthlyCompleteStatusUpdateParams) (MonthlyGoal, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UserDailyGoals(ctx context.Context, userID sql.NullInt32) ([]DailyGoal, error)
	UserMonthlyCompletedGoals(ctx context.Context, userID sql.NullInt32) ([]MonthlyGoal, error)
	UserMonthlyGoals(ctx context.Context, userID sql.NullInt32) ([]MonthlyGoal, error)
	UserMonthlyUncompletedGoals(ctx context.Context, userID sql.NullInt32) ([]MonthlyGoal, error)
	UserWeeklyGoals(ctx context.Context, userID sql.NullInt32) ([]WeeklyGoal, error)
	WeeklyCompleteStatusUpdate(ctx context.Context, arg WeeklyCompleteStatusUpdateParams) (WeeklyGoal, error)
	WeeklyCompletedGoals(ctx context.Context, userID sql.NullInt32) ([]WeeklyGoal, error)
	WeeklyUncompletedGoals(ctx context.Context, userID sql.NullInt32) ([]WeeklyGoal, error)
}

var _ Querier = (*Queries)(nil)
