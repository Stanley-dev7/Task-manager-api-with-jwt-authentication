package main

import (
	"log"
	"net/http"

	"day34-task-manager-api/database"
	"day34-task-manager-api/handlers"
	"day34-task-manager-api/middleware"

	"github.com/gorilla/mux"
)

func main() {

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

