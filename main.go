package main

import (
	"log"
	"net/http"
)

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Привет!"))
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Если это не так, то вызывается метод w.WriteHeader() для возвращения статус-кода 405
		// и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
		// Затем мы завершаем работу функции вызвав "return", чтобы
		// последующий код не выполнялся.
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!"))
		return
	}
	w.Write([]byte("Форма для создания новой заметки..."))
}

func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
