import (
"fmt"
"log"
"net/http"
)

const Port = "8880"

func main() {
// 1. Раздача статических файлов (CSS, картинки)
// Если у вас есть файл style.css в папке static, он будет доступен по адресу /static/style.css
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

// 2. Раздача HTML шаблонов
// Ваш файл index.html должен лежать в папке templates
http.HandleFunc("/", indexHandler)

// 3. API эндпоинты (как на скриншоте server.go)
http.HandleFunc("/api/register", registerHandler)
http.HandleFunc("/api/login", loginHandler)
http.HandleFunc("/api/cart", cartHandler) // Допустим, для корзины

fmt.Printf("Сервер запущен: http://localhost:%s\n", Port)
log.Fatal(http.ListenAndServe(":"+Port, nil))
}
// Говорим Go: "Это главный файл, с которого всё начинается"
package main
// подключаем инструменты
import (
    "fmt" //  для печати текста в консоль(информация о действиях на сайте)
    "net/http" // для создания веб-сервера    
    "log"
    "html/template"
)
func main() { // главная функция, которая запускается автоматически
    http.HandleFunc("/", indexHandler) // "Когда кто-то заходит на сайт (на главную страницу /), делай следующее..."


        // Запускаем сервер
	fmt.Println("Сервер запущен на http://localhost:8080") // Сервер работает, заходи по этому адресу
	fmt.Println("Нажмите Ctrl+C для остановки")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Запускаем сервер и говорим ему слушать порт 8080('номер квартиры')
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	//Любой путь кроме "/" отправляем на 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}
