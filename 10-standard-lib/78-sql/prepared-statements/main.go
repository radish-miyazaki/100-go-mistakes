package main

import "database/sql"

func listing(db *sql.DB, id string) error {
	stmt, err := db.Prepare("SELECT * FROM order WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return err
	}

	_ = rows
	return nil
}
