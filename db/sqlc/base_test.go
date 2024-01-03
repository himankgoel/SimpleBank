package db

import (
	"database/sql"
	"log"
	"testing"
	"os"
	
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Connection to db failed!")
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
