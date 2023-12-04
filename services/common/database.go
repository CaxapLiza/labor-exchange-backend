package common

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase() (*Database, error) {
	connectionString := "user=postgres password=qwerty dbname=labor_exchange_db sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to the database")

	return &Database{Connection: db}, nil
}

func (db *Database) Close() {
	db.Connection.Close()
}
