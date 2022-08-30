package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/junimslage10/gofinance-backend/api"
	db "github.com/junimslage10/gofinance-backend/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:root@localhost:5432/go_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api: ", err)
	}
}
