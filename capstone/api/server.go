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

func SetUpRouter() *gin.Engine {
	dbInit()
	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
	r.DELETE("/users/:id", deleteUser)
	r.PUT("/users/:id", updateUser)

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
