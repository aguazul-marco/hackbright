package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/aguazul-marco/hackbright/capstone/db/sqlc"
	"github.com/gin-gonic/gin"
)

func createMonthlyGoal(ctx *gin.Context) {
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

	args := db.CreateMonthlyGoalParams{
		Discription: req.Discription,
		UserID:      userid,
	}

	goal, err := q.CreateMonthlyGoal(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goal)
}

func updateMonthlyGoalStatus(ctx *gin.Context) {
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

	args := db.MonthlyCompleteStatusUpdateParams{
		ID:        goalID,
		Completed: req.Completed,
	}

	goalUpdate, err := q.MonthlyCompleteStatusUpdate(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goalUpdate)

}

func getMonthlyCompletedGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	completedGoals, err := q.UserMonthlyCompletedGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, completedGoals)
}

func getMonthlyUncompletedGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	uncompletedGoals, err := q.UserMonthlyUncompletedGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, uncompletedGoals)
}

func getAllMonthlyGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	goals, err := q.UserMonthlyGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.JSON(http.StatusOK, goals)
}

func deleteMonthlyGoal(ctx *gin.Context) {
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

	err = q.DeleteMonthlyGoal(ctx, goalID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.Status(http.StatusOK)

}
