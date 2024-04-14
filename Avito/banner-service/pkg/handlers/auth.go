package handlers

import (
	"banner-service/pkg/models"
	"encoding/json"
	"net/http"
	"time"
)

// Signin обрабатывает вход пользователя и генерирует JWT токен
func Signin(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	isAdmin := false
	if creds.Username == "admin" && creds.Password == "password123" {
		isAdmin = true
	}

	// Генерация JWT токена
	role := "user"
	if isAdmin {
		role = "admin"
	}
	token, err := models.GenerateJWT(creds.Username, role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
