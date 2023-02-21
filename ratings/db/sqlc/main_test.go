package db

import (
	"database/sql"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql:///test_ratings?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	// Finds test_ratings and dropping if exists
	dropdb := exec.Command("dropdb", "--if-exists", "test_ratings")
	if err := dropdb.Run(); err != nil {
		log.Fatal(err)
	}

	// Opens connection to test_ratings
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	mi, err := migrate.NewWithDatabaseInstance(
		"file:///db/migrations",
		"postgres", driver)
	mi.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	// Question: why does db work here? 👇
	testQueries = New(db)

	os.Exit(m.Run())
}
