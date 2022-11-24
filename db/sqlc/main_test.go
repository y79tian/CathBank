package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_"github.com/lib/pq" // we do not directly use it, must add _ first, need this bc database/sql is only geenric interface
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:3001/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

// main entry point
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}
	testQueries = New(testDB)
	// m.Run() to run main tests, return an exit code
	os.Exit(m.Run())
}

