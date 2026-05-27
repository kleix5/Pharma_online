package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Структура для ответа API
type ResponseType struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// 1. Обработчик главной страницы (index.html)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Отдаем файл из папки templates
	http.ServeFile(w, r, "index.html")
}

// 2. Обработчик регистрации
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

	// Здесь должна быть логика сохранения в БД
	log.Printf("Пользователь зарегистрирован: %s", email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(ResponseType{
		Status:  "success",
		Message: "User registered",
	})
}

// 3. Обработчик логина
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

	// Имитация проверки пароля
	if req.Email == "user@test.com" && req.Password == "pass" {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(ResponseType{
			Status: "success",
			Data: map[string]string{
				"token": "fake-jwt-token",
			},
		})

		return
	}

	writeError(w, http.StatusUnauthorized, "Invalid credentials")
}

// 4. Обработчик корзины
func cartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Возвращаем пустую корзину для примера
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(ResponseType{
			Status: "success",
			Data:   []string{},
		})

		return
	}

	writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

// Вспомогательная функция для отправки ошибок
func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(ResponseType{
		Status:  "error",
		Message: msg,
	})
}
