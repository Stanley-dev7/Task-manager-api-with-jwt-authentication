package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"day34-task-manager-api/database"
	"day34-task-manager-api/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

// 🔐 JWT secret key (for now hardcoded, later we move to env)
var jwtKey = []byte("secret_key_123")

// =====================
// REGISTER USER
// =====================
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save user
	_, err = database.DB.Exec(
		"INSERT INTO users(email, password) VALUES(?, ?)",
		user.Email,
		string(hashedPassword),
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("User registered successfully")
}

// =====================
// LOGIN USER + JWT
// =====================
func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var storedUser models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user from DB
	err = database.DB.QueryRow(
		"SELECT id, email, password FROM users WHERE email = ?",
		user.Email,
	).Scan(&storedUser.ID, &storedUser.Email, &storedUser.Password)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword(
		[]byte(storedUser.Password),
		[]byte(user.Password),
	)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Return token
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}