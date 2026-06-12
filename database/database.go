package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() error {
	var err error

	DB, err = sql.Open("sqlite", "todos.db")
	if err != nil {
		return err
	}

	todoQuery := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE,
		user_id TEXT
	);
	`

	userQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`

	_, err = DB.Exec(todoQuery)
	if err != nil {
		return err
	}

	_, err = DB.Exec(userQuery)
	if err != nil {
		return err
	}

	return nil
}