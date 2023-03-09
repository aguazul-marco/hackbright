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

	r.POST("/user", createUser)
	r.POST("/user/:id/dailygoal", createGoal)
	r.GET("/user/:id", getUser)
	r.GET("/user/:id/dailygoals", getAllDailyGoals)
	r.GET("/user/:id/dailygoals/completed", getCompleteGoals)
	r.GET("/user/:id/dailygoals/uncompleted", getUnCompleteGoals)
	r.DELETE("user/:id/dailygoal/:goal_id", deleteGoal)
	r.DELETE("/user/:id", deleteUser)
	r.PUT("/user/:id", updateUser)
	r.PUT("/user/:id/dailygoal/:goal_id", updateGoalStatus)

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
