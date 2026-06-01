package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ===== CONFIG =====

const Port = "8080"

// ===== STRUCTS =====

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ===== MAIN =====

func main() {
	// Статика
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// HTML
	http.HandleFunc("/", indexHandler)

	// API
	http.HandleFunc("/api/register", registerHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/cart", cartHandler)

	fmt.Printf("Сервер запущен: http://localhost:%s\n", Port)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}

// ===== HANDLERS =====

// Главная страница
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

// Регистрация
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	password := strings.TrimSpace(req.Password)

	if email == "" || password == "" {
		writeError(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	log.Printf("Пользователь зарегистрирован: %s", email)

	writeJSON(w, http.StatusCreated, Response{
		Status:  "success",
		Message: "User registered",
	})
}

// Логин
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Email == "user@test.com" && req.Password == "pass" {
		writeJSON(w, http.StatusOK, Response{
			Status: "success",
			Data:   map[string]string{"token": "fake-jwt-token"},
		})
		return
	}

	writeError(w, http.StatusUnauthorized, "Invalid credentials")
}

// Корзина
func cartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writeJSON(w, http.StatusOK, Response{
			Status: "success",
			Data:   []string{},
		})
		return
	}

	writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

// ===== HELPERS =====

func writeError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, Response{
		Status:  "error",
		Message: msg,
	})
}

func writeJSON(w http.ResponseWriter, code int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}
