package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/aguazul-marco/hackbright/capstone/db/sqlc"
	"github.com/gin-gonic/gin"
)

func createWeeklyGoal(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	var req CreateGoalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	args := db.CreateWeeklyGoalParams{
		Discription: req.Discription,
		UserID:      userid,
	}

	goal, err := q.CreateWeeklyGoal(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goal)
}

func updateWeeklyGoalStatus(ctx *gin.Context) {
	uID := ctx.Param("id")
	_, err := strconv.ParseInt(uID, 10, 64)
	if err != nil {
		fmt.Println("error occured converting user ID:", err)
	}

	id := ctx.Param("goal_id")
	goalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting goal ID:", err)
	}

	var req GoalStatusUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	args := db.WeeklyCompleteStatusUpdateParams{
		ID:        goalID,
		Completed: req.Completed,
	}

	goalUpdate, err := q.WeeklyCompleteStatusUpdate(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goalUpdate)

}

func getWeeklyCompletedGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	completedGoals, err := q.WeeklyCompletedGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, completedGoals)
}

func getWeeklyUncompletedGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	uncompletedGoals, err := q.WeeklyUncompletedGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, uncompletedGoals)
}

func getAllWeeklyGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	goals, err := q.UserWeeklyGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.JSON(http.StatusOK, goals)
}

func deleteWeeklyGoal(ctx *gin.Context) {
	uID := ctx.Param("id")
	_, err := strconv.ParseInt(uID, 10, 64)
	if err != nil {
		fmt.Println("error occured converting user ID:", err)
	}

	id := ctx.Param("goal_id")
	goalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting goal ID:", err)
	}

	err = q.DeleteWeeklyGoal(ctx, goalID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.Status(http.StatusOK)

}
