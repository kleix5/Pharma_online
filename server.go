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
