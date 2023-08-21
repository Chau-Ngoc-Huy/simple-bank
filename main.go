package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/simple-bank/api"
	db "github.com/simple-bank/db/sqlc"
	"github.com/simple-bank/util"
	"log"
)

//const (
//	dbDriver      = "postgres"
//	dbSource      = "postgresql://root:12345678@localhost:5432/simple_bank?sslmode=disable"
//	serverAddress = "0.0.0.0:8080"
//)

func main() {

	conf, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("Can not load config")
	}
	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(conf.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
