package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-learning/day32-task-manager-db/database"
	"go-learning/day32-task-manager-db/models"
)

/* -----------------------------
   GET /todos & POST /todos
------------------------------*/

func TodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	// GET ALL TODOS (ONLY FOR LOGGED IN USER)
	case "GET":

		email := r.Header.Get("user_email")

		rows, err := database.DB.Query(
			"SELECT id, title, completed FROM todos WHERE user_id = ?",
			email,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var todos []models.Todo

		for rows.Next() {
			var todo models.Todo

			rows.Scan(
				&todo.ID,
				&todo.Title,
				&todo.Completed,
			)

			todos = append(todos, todo)
		}

		json.NewEncoder(w).Encode(todos)

	// CREATE TODO (ASSIGNED TO LOGGED IN USER)
	case "POST":

		var todo models.Todo

		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		email := r.Header.Get("user_email")

		_, err = database.DB.Exec(
			"INSERT INTO todos(title, completed, user_id) VALUES(?, ?, ?)",
			todo.Title,
			false,
			email,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Todo created successfully",
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/* -----------------------------
   GET / PUT / DELETE /todos/{id}
------------------------------*/

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case "GET":

		var todo models.Todo

		err := database.DB.QueryRow(
			"SELECT id, title, completed FROM todos WHERE id = ?",
			id,
		).Scan(&todo.ID, &todo.Title, &todo.Completed)

		if err != nil {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(todo)

	case "PUT":

		var todo models.Todo

		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		_, err = database.DB.Exec(
			"UPDATE todos SET title = ?, completed = ? WHERE id = ?",
			todo.Title,
			todo.Completed,
			id,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Todo updated successfully",
		})

	case "DELETE":

		_, err := database.DB.Exec(
			"DELETE FROM todos WHERE id = ?",
			id,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Todo deleted successfully",
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}