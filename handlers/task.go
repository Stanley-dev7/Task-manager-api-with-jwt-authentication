package handlers

import (
	"encoding/json"
	"net/http"

	"day34-task-manager-api/database"
	"day34-task-manager-api/models"
	"day34-task-manager-api/middleware"
	"github.com/gorilla/mux"
)

// =====================
// CREATE TASK (SECURE)
// =====================
func CreateTask(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user_id from JWT context (NOT from request body)
	userID := r.Context().Value(middleware.UserIDKey).(int)

	_, err = database.DB.Exec(
		"INSERT INTO tasks(title, description, status, priority, due_date, user_id) VALUES(?,?,?,?,?,?)",
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.DueDate,
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Task created successfully")
}

// =====================
// GET TASKS (ONLY USER TASKS)
// =====================
func GetTasks(w http.ResponseWriter, r *http.Request) {

	// Get user_id from JWT context
	userID := r.Context().Value(middleware.UserIDKey).(int)

	rows, err := database.DB.Query(
		"SELECT id, title, description, status, priority, due_date, user_id FROM tasks WHERE user_id = ?",
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {

		var task models.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.DueDate,
			&task.UserID,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	json.NewEncoder(w).Encode(tasks)
}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int)

	id := mux.Vars(r)["id"]

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"UPDATE tasks SET title=?, description=?, status=?, priority=?, due_date=? WHERE id=? AND user_id=?",
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.DueDate,
		id,
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Task updated successfully")
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int)

	id := mux.Vars(r)["id"]

	_, err := database.DB.Exec(
		"DELETE FROM tasks WHERE id=? AND user_id=?",
		id,
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Task deleted successfully")
}