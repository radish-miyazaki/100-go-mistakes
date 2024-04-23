package main

import (
	"database/sql"
	"errors"
	"log"
)

const query = "..."

func getBalance1(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	return 0, nil
}

func getBalance2(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() { _ = rows.Close() }()

	return 0, nil
}

func getBalance3(db *sql.DB, clientID string) (balance float32, err error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		if err != nil {
			if closeErr != nil {
				log.Printf("failed to close rows: %w", closeErr)
			}

			return
		}

		err = closeErr
	}()

	if rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			return 0, err
		}

		return balance, nil
	}

	return 0, nil
}

func getBalance4(db *sql.DB, clientID string) (balance float32, err error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}

	var scanErr error
	defer func() {
		closeErr := rows.Close()
		if scanErr != nil && closeErr != nil {
			err = errors.Join(scanErr, closeErr)
			return
		}

		if scanErr != nil {
			err = scanErr
			return
		}

		if closeErr != nil {
			err = closeErr
			return
		}
	}()

	if rows.Next() {
		var balance float32
		scanErr := rows.Scan(&balance)
		if scanErr != nil {
			return 0, scanErr
		}

		return balance, nil
	}

	return 0, nil
}
