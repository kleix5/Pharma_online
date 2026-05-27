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
