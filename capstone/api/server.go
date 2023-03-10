package api

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/aguazul-marco/hackbright/capstone/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql:///goal_tracker?sslmode=disable"
)

var q *db.Queries

func SetUpRouter() *gin.Engine {
	dbInit()
	r := gin.Default()

	r.POST("/user/", createUser)
	r.GET("/user/:id", getUser)
	r.PUT("/user/:id", updateUser)
	r.DELETE("/user/:id", deleteUser)

	//daily goals endpoints
	r.POST("/user/:id/dailygoal", createDailyGoal)
	r.GET("/user/:id/dailygoals", getAllDailyGoals)
	r.GET("/user/:id/dailygoals/completed", getDailyCompletedGoals)
	r.GET("/user/:id/dailygoals/uncompleted", getDailyUncompletedGoals)
	r.PUT("/user/:id/dailygoal/:goal_id", updateDailyGoalStatus)
	r.DELETE("user/:id/dailygoal/:goal_id", deleteDailyGoal)

	//monthly goals endpoints
	r.POST("/user/:id/monthlygoal", createMonthlyGoal)
	r.GET("/user/:id/monthlygoals", getAllMonthlyGoals)
	r.GET("/user/:id/monthlygoals/completed", getMonthlyCompletedGoals)
	r.GET("/user/:id/monthlygoals/uncompleted", getMonthlyUncompletedGoals)
	r.PUT("/user/:id/monthlygoal/:goal_id", updateMonthlyGoalStatus)
	r.DELETE("user/:id/monthlygoal/:goal_id", deleteMonthlyGoal)

	//weekly goals endpoints
	r.POST("/user/:id/weeklygoal", createWeeklyGoal)
	r.GET("/user/:id/weeklygoals", getAllWeeklyGoals)
	r.GET("/user/:id/weeklygoals/completed", getWeeklyCompletedGoals)
	r.GET("/user/:id/weeklygoals/uncompleted", getWeeklyUncompletedGoals)
	r.PUT("/user/:id/weeklygoal/:goal_id", updateWeeklyGoalStatus)
	r.DELETE("user/:id/weeklygoal/:goal_id", deleteWeeklyGoal)

	return r
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func dbInit() {
	database, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatalf("db is not connected: %v", err)
	} else {
		fmt.Println("db is connected")
	}

	q = db.New(database)
}
