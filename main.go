package main

import (
	"log"
	"net/http"
)

func main() {
	// Главная страница (index.html)
	http.Handle("/", http.FileServer(http.Dir("./")))

	// Папка с товарами (поддерживает русские имена)
	http.Handle(
		"/products/",
		http.StripPrefix(
			"/products/",
			http.FileServer(http.Dir("./products")),
		),
	)

	log.Println("Сервер запущен: http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
