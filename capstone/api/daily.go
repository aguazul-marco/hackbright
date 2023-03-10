package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/aguazul-marco/hackbright/capstone/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateGoalRequest struct {
	Discription string        `json:"discription" binding:"required"`
	UserID      sql.NullInt32 `json:"user_id"`
}

type GoalStatusUpdateRequest struct {
	ID        int64 `json:"goal_id"`
	Completed bool  `json:"completed"`
}

func createDailyGoal(ctx *gin.Context) {
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

func updateDailyGoalStatus(ctx *gin.Context) {
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

func getDailyCompletedGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}

	if err := validate(i); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	completedGoals, err := q.DailyCompletedGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	if err := validate(completedGoals); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, completedGoals)
}

func getDailyUncompletedGoals(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("error occured converting:", err)
	}

	if err := validate(i); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	uncompletedGoals, err := q.DailyUncompletedGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	if err := validate(uncompletedGoals); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
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

	if err := validate(i); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	userid := sql.NullInt32{Int32: int32(i), Valid: true}

	goals, err := q.UserDailyGoals(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	if err := validate(goals); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goals)
}

func deleteDailyGoal(ctx *gin.Context) {
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

	if err := validate(goalID); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err = q.DeleteDailyGoal(ctx, goalID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.JSON(http.StatusOK, "Daily goal deleted")
}

func validate(i interface{}) (err error) {
	if i == 0 {
		err = errors.New("id doesn't exist: try again")
	}
	if i == "" {
		err = errors.New("empty value: please enter a goal")
	}
	if i == nil {
		err = errors.New("no goals")
	}
	return err
}
