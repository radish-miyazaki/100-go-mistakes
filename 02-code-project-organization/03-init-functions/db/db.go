package db

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	dsn := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}

	db = d
}

func createClient(dsn string) (*sql.DB, error) {
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = d.Ping()
	if err != nil {
		return nil, err
	}

	return d, nil
}
