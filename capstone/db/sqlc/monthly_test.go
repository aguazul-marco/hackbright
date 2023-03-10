package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/aguazul-marco/hackbright/capstone/util"
	"github.com/stretchr/testify/require"
)

func createRandomMonthlyGoal(t *testing.T) MonthlyGoal {
	user1 := createRandomUser(t)
	userid := sql.NullInt32{Int32: int32(user1.ID), Valid: true}

	args1 := CreateMonthlyGoalParams{
		Discription: util.RandomString(15),
		UserID:      userid,
	}

	goal, err := testQueries.CreateMonthlyGoal(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goal)

	require.Equal(t, goal.UserID, userid)
	require.Equal(t, goal.Discription, args1.Discription)
	require.False(t, false, goal.Completed)

	require.NotZero(t, goal.UserID)
	require.NotZero(t, goal.CreatedAt)

	return goal
}

func TestCreatMonthlyGoal(t *testing.T) {
	createRandomMonthlyGoal(t)
}

func TestMonthlyCompleteStatusUpdate(t *testing.T) {
	goal1 := createRandomMonthlyGoal(t)

	args1 := MonthlyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: true,
	}

	goalComplete, err := testQueries.MonthlyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalComplete)

	require.True(t, true, goalComplete)
	require.Equal(t, goal1.ID, goalComplete.ID)
	require.Equal(t, args1.ID, goalComplete.ID)
	require.Equal(t, args1.Completed, goalComplete.Completed)

	args2 := MonthlyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: false,
	}

	goalUncomplete, err := testQueries.MonthlyCompleteStatusUpdate(context.Background(), args2)
	require.NoError(t, err)
	require.NotEmpty(t, goalUncomplete)

	require.False(t, false, goalUncomplete)
	require.Equal(t, goal1.ID, goalUncomplete.ID)
	require.Equal(t, args2.ID, goalUncomplete.ID)
	require.Equal(t, args2.Completed, goalUncomplete.Completed)

}

func TestMonthlyCompletedGoals(t *testing.T) {
	goal1 := createRandomMonthlyGoal(t)

	args1 := MonthlyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: true,
	}

	goalComplete, err := testQueries.MonthlyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalComplete)

	completedGoals, err := testQueries.UserMonthlyCompletedGoals(context.Background(), goalComplete.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, completedGoals)

	require.Equal(t, goalComplete.Completed, args1.Completed)

}

func TestMonthlyUncompletedGoals(t *testing.T) {
	goal1 := createRandomMonthlyGoal(t)

	args1 := MonthlyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: false,
	}

	goalUncomplete, err := testQueries.MonthlyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalUncomplete)

	uncompletedGoals, err := testQueries.UserMonthlyUncompletedGoals(context.Background(), goalUncomplete.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, uncompletedGoals)

	require.Equal(t, goalUncomplete.Completed, args1.Completed)

}

func TestDeleteMonthlyGoal(t *testing.T) {
	goal1 := createRandomMonthlyGoal(t)

	err := testQueries.DeleteMonthlyGoal(context.Background(), goal1.ID)
	require.NoError(t, err)

}

func TestUserMonthlyGoals(t *testing.T) {
	var g MonthlyGoal
	for i := 0; i < 5; i++ {
		g = createRandomMonthlyGoal(t)
	}

	goals, err := testQueries.UserMonthlyGoals(context.Background(), g.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, goals)

	for _, goal := range goals {
		require.NotEmpty(t, goal)
	}
}
