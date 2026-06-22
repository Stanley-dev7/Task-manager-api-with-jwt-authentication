package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

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
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`

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
}