package database

import (
	"database/sql"
<<<<<<< HEAD
	"log"
=======
>>>>>>> 4995461d3620718635e8e7a3bbcb0bfbf828ca9e

	_ "modernc.org/sqlite"
)

var DB *sql.DB

<<<<<<< HEAD
func Connect() {
	var err error

	DB, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully")
}

func InitTables() {

	userTable := `
=======
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
>>>>>>> 4995461d3620718635e8e7a3bbcb0bfbf828ca9e
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`

<<<<<<< HEAD
	taskTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending',
		priority TEXT DEFAULT 'medium',
		due_date TEXT,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(taskTable)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tables created successfully")
=======
	_, err = DB.Exec(todoQuery)
	if err != nil {
		return err
	}

	_, err = DB.Exec(userQuery)
	if err != nil {
		return err
	}

	return nil
>>>>>>> 4995461d3620718635e8e7a3bbcb0bfbf828ca9e
}