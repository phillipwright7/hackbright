package db

import (
	"database/sql"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql:///test_ratings?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	cmd := exec.Command("dropdb", "--if-exists", "test_ratings")
	if err := cmd.Run(); err != nil {
		log.Fatalf("dropdb failed: %v", err)
	}

	cmd = exec.Command("createdb", "test_ratings")
	if err := cmd.Run(); err != nil {
		log.Fatalf("createdb failed: %v", err)
	}

	mi, err := migrate.New(
		"file://../migrations",
		dataSourceName)
	if err != nil {
		log.Fatalf("failed to create new migrate instance: %v", err)
	}

	if err = mi.Up(); err != nil {
		log.Fatalf("failed to migrate up: %v", err)
	}

	testdb, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	testQueries = New(testdb)

	os.Exit(m.Run())
}
