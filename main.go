package main

import (
<<<<<<< HEAD
	"log"
	"net/http"

	"day34-task-manager-api/database"
	"day34-task-manager-api/handlers"
	"day34-task-manager-api/middleware"

	"github.com/gorilla/mux"
=======
	"fmt"
	"log"
	"net/http"

	"go-learning/day32-task-manager-db/database"
	"go-learning/day32-task-manager-db/handlers"
>>>>>>> 4995461d3620718635e8e7a3bbcb0bfbf828ca9e
)

func main() {

<<<<<<< HEAD
	database.Connect()
	database.InitTables()

	r := mux.NewRouter()

	// AUTH ROUTES
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// PROTECTED ROUTES
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.ValidateToken)

	api.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	api.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	api.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("Put")
	api.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("Delete")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

=======
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
>>>>>>> 4995461d3620718635e8e7a3bbcb0bfbf828ca9e
