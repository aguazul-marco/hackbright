package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql:///test_ratings?sslmode=disable"
)

func TestMain(m *testing.M) {
	cmd := exec.Command("dropdb", "--if-exists", "test_ratings")
	if err := cmd.Run(); err != nil {
		fmt.Println("failed to drop")
		log.Fatal(err)
	}

	createdb := exec.Command("createdb", "test_ratings")
	if err := createdb.Run(); err != nil {
		log.Fatalf("failed to create: %v", err)
	}

	mig, err := migrate.New(
		"file://../migration",
		dbSource)
	if err != nil {
		log.Fatal(err)
	}
	if err := mig.Up(); err != nil {
		log.Fatal(err)
	}

	connect, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(connect)

	os.Exit(m.Run())

}
