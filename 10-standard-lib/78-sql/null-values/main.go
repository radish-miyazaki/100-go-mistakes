package main

import "database/sql"

func listing1(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT dep, age FROM emp WHERE id = ?", id)
	if err != nil {
		return err
	}

	var (
		department string
		age        int
	)

	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}

	return nil
}

func listing2(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT dep, age FROM emp WHERE id = ?", id)
	if err != nil {
		return err
	}

	var (
		department *string
		age        int
	)

	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}

	return nil
}

func listing3(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT dep, age FROM emp WHERE id = ?", id)
	if err != nil {
		return err
	}

	var (
		department sql.NullString
		age        int
	)

	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}

	return nil
}
