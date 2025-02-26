package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/dhvll/go-advance/api"
	db "github.com/dhvll/go-advance/db/sqlc"
	"github.com/dhvll/go-advance/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect to config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
