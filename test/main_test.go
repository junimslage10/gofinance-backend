package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/junimslage10/gofinance-backend/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:root@localhost:5432/go_finance?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = db.New(conn)
	os.Exit(m.Run())
}
