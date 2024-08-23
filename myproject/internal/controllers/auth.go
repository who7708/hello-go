package controllers

import (
	"net/http"

	"myproject/internal/services"
	"myproject/pkg/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	err := services.RegisterUser(username, password)
	if err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	success, err := services.LoginUser(username, password)
	if err != nil {
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}
	if success {
		token, err := utils.GenerateToken(username)
		if err != nil {
			http.Error(w, "Token generation failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
