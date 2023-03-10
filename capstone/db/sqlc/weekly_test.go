package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/aguazul-marco/hackbright/capstone/util"
	"github.com/stretchr/testify/require"
)

func createRandomWeeklyGoal(t *testing.T) WeeklyGoal {
	user1 := createRandomUser(t)
	userid := sql.NullInt32{Int32: int32(user1.ID), Valid: true}

	args1 := CreateWeeklyGoalParams{
		Discription: util.RandomString(15),
		UserID:      userid,
	}

	goal, err := testQueries.CreateWeeklyGoal(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goal)

	require.Equal(t, goal.UserID, userid)
	require.Equal(t, goal.Discription, args1.Discription)
	require.False(t, false, goal.Completed)

	require.NotZero(t, goal.UserID)
	require.NotZero(t, goal.CreatedAt)

	return goal
}

func TestCreatWeeklyGoal(t *testing.T) {
	createRandomWeeklyGoal(t)
}

func TestWeeklyCompleteStatusUpdate(t *testing.T) {
	goal1 := createRandomWeeklyGoal(t)

	args1 := WeeklyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: true,
	}

	goalComplete, err := testQueries.WeeklyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalComplete)

	require.True(t, true, goalComplete)
	require.Equal(t, goal1.ID, goalComplete.ID)
	require.Equal(t, args1.ID, goalComplete.ID)
	require.Equal(t, args1.Completed, goalComplete.Completed)

	args2 := WeeklyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: false,
	}

	goalUncomplete, err := testQueries.WeeklyCompleteStatusUpdate(context.Background(), args2)
	require.NoError(t, err)
	require.NotEmpty(t, goalUncomplete)

	require.False(t, false, goalUncomplete)
	require.Equal(t, goal1.ID, goalUncomplete.ID)
	require.Equal(t, args2.ID, goalUncomplete.ID)
	require.Equal(t, args2.Completed, goalUncomplete.Completed)

}

func TestWeeklyCompletedGoals(t *testing.T) {
	goal1 := createRandomWeeklyGoal(t)

	args1 := WeeklyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: true,
	}

	goalComplete, err := testQueries.WeeklyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalComplete)

	completedGoals, err := testQueries.WeeklyCompletedGoals(context.Background(), goalComplete.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, completedGoals)

	require.Equal(t, goalComplete.Completed, args1.Completed)

}

func TestWeeklyUncompletedGoals(t *testing.T) {
	goal1 := createRandomWeeklyGoal(t)

	args1 := WeeklyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: false,
	}

	goalUncomplete, err := testQueries.WeeklyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalUncomplete)

	uncompletedGoals, err := testQueries.WeeklyUncompletedGoals(context.Background(), goalUncomplete.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, uncompletedGoals)

	require.Equal(t, goalUncomplete.Completed, args1.Completed)

}

func TestDeleteWeeklyGoal(t *testing.T) {
	goal1 := createRandomWeeklyGoal(t)

	err := testQueries.DeleteWeeklyGoal(context.Background(), goal1.ID)
	require.NoError(t, err)

}

func TestUserWeeklyGoals(t *testing.T) {
	var g WeeklyGoal
	for i := 0; i < 5; i++ {
		g = createRandomWeeklyGoal(t)
	}

	goals, err := testQueries.UserWeeklyGoals(context.Background(), g.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, goals)

	for _, goal := range goals {
		require.NotEmpty(t, goal)
	}
}
