package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // we do not directly use it, must add _ first, need this bc database/sql is only geenric interface
	"github.com/techschool/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

// main entry point
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}
	testQueries = New(testDB)
	// m.Run() to run main tests, return an exit code
	os.Exit(m.Run())
}

