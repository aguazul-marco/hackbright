package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/aguazul-marco/hackbright/capstone/util"
	"github.com/stretchr/testify/require"
)

func createRandomDailyGoal(t *testing.T) DailyGoal {
	user1 := createRandomUser(t)
	userid := sql.NullInt32{Int32: int32(user1.ID), Valid: true}

	args1 := CreateDailyGoalParams{
		Discription: util.RandomString(15),
		UserID:      userid,
	}

	goal, err := testQueries.CreateDailyGoal(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goal)

	require.Equal(t, goal.UserID, userid)
	require.Equal(t, goal.Discription, args1.Discription)
	require.False(t, false, goal.Completed)

	require.NotZero(t, goal.UserID)
	require.NotZero(t, goal.CreatedAt)

	return goal
}

func TestCreatDailyGoal(t *testing.T) {
	createRandomDailyGoal(t)
}

func TestDailyCompleteStatusUpdate(t *testing.T) {
	goal1 := createRandomDailyGoal(t)

	args1 := DailyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: true,
	}

	goalComplete, err := testQueries.DailyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalComplete)

	require.True(t, true, goalComplete)
	require.Equal(t, goal1.ID, goalComplete.ID)
	require.Equal(t, args1.ID, goalComplete.ID)
	require.Equal(t, args1.Completed, goalComplete.Completed)

	args2 := DailyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: false,
	}

	goalUncomplete, err := testQueries.DailyCompleteStatusUpdate(context.Background(), args2)
	require.NoError(t, err)
	require.NotEmpty(t, goalUncomplete)

	require.False(t, false, goalUncomplete)
	require.Equal(t, goal1.ID, goalUncomplete.ID)
	require.Equal(t, args2.ID, goalUncomplete.ID)
	require.Equal(t, args2.Completed, goalUncomplete.Completed)

}

func TestDailyCompletedGoals(t *testing.T) {
	goal1 := createRandomDailyGoal(t)

	args1 := DailyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: true,
	}

	goalComplete, err := testQueries.DailyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalComplete)

	completedGoals, err := testQueries.DailyCompletedGoals(context.Background(), goalComplete.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, completedGoals)

	require.Equal(t, goalComplete.Completed, args1.Completed)

}

func TestDailyUncompletedGoals(t *testing.T) {
	goal1 := createRandomDailyGoal(t)

	args1 := DailyCompleteStatusUpdateParams{
		ID:        goal1.ID,
		Completed: false,
	}

	goalUncomplete, err := testQueries.DailyCompleteStatusUpdate(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, goalUncomplete)

	uncompletedGoals, err := testQueries.DailyUncompletedGoals(context.Background(), goalUncomplete.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, uncompletedGoals)

	require.Equal(t, goalUncomplete.Completed, args1.Completed)

}

func TestDeleteDailyGoal(t *testing.T) {
	goal1 := createRandomDailyGoal(t)

	err := testQueries.DeleteDailyGoal(context.Background(), goal1.ID)
	require.NoError(t, err)

}

func TestUserDailyGoals(t *testing.T) {
	var g DailyGoal
	for i := 0; i < 5; i++ {
		g = createRandomDailyGoal(t)
	}

	goals, err := testQueries.UserDailyGoals(context.Background(), g.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, goals)

	for _, goal := range goals {
		require.NotEmpty(t, goal)
	}
}
