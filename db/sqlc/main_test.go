package db

import (
	"database/sql"
	"github.com/simple-bank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

//const (
//	dbDriver = "postgres"
//	dbSource = "postgresql://root:12345678@localhost:5432/simple_bank?sslmode=disable"
//)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	conf, err := util.LoadConfig("./../../")
	if err != nil {
		log.Fatal("can not load config", err)
	}

	testDB, err = sql.Open(conf.DBDriver, conf.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
