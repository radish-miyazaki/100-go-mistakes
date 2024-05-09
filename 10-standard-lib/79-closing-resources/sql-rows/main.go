package main

import (
	"database/sql"
	"log"
)

var dataSourceName string

func listing1() error {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		return err
	}

	_ = rows

	return nil
}

func listing2() error {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalf("failed to close rows: %v\n", err)
		}
	}()

	_ = rows

	return nil
}
