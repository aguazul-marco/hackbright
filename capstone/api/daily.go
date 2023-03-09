package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/aguazul-marco/hackbright/capstone/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateDailyGoalRequest struct {
	Discription string        `json:"discription" binding:"required"`
	UserID      sql.NullInt32 `json:"user_id"`
}

type GoalStatusUpdateRequest struct {
	ID        int64 `json:"goal_id"`
	Completed bool  `json:"completed"`
}

func createGoal(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	var req CreateDailyGoalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	args := db.CreateDailyGoalParams{
		Discription: req.Discription,
		UserID:      userid,
	}

	goal, err := q.CreateDailyGoal(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goal)

}

func updateGoalStatus(ctx *gin.Context) {
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

	args := db.DailyCompleteStatusUpdateParams{
		ID:        goalID,
		Completed: req.Completed,
	}

	goalUpdate, err := q.DailyCompleteStatusUpdate(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goalUpdate)

}

func getCompleteGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	completedGoals, err := q.DailyCompletedGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, completedGoals)
}

func getUnCompleteGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	uncompletedGoals, err := q.DailyUncompletedGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, uncompletedGoals)
}

func getAllDailyGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}
	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	goals, err := q.UserDailyGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.JSON(http.StatusOK, goals)

}

func deleteGoal(ctx *gin.Context) {
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

	err = q.DeleteDailyGoal(ctx, goalID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.Status(http.StatusOK)

}
