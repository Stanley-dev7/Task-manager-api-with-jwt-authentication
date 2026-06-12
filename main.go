package main

import (
	"fmt"
	"log"
	"net/http"

	"go-learning/day32-task-manager-db/database"
	"go-learning/day32-task-manager-db/handlers"
)

func main() {

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully!")

	http.HandleFunc("/todos", handlers.AuthMiddleware(handlers.TodosHandler))
http.HandleFunc("/todos/", handlers.AuthMiddleware(handlers.TodoHandler))
	http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
	fmt.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}