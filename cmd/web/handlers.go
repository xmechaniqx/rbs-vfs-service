package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
) // Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Привет из Snippetbox" как тело ответа.

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/index.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	// w.Write([]byte("Привет из Snippetbox (handler)"))
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Используем функцию fmt.Fprintf() для вставки значения из id в строку ответа
	// и записываем его в http.ResponseWriter.
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Если это не так, то вызывается метод w.WriteHeader() для возвращения статус-кода 405
		// и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
		// Затем мы завершаем работу функции вызвав "return", чтобы
		// последующий код не выполнялся.
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		http.Error(w, "Метод запрещен!", 405)
		// w.Write([]byte("GET-Метод запрещен!"))
		return
	}
	w.Write([]byte("Форма для создания новой заметки..."))
}

// func showFlag(w http.ResponseWriter, r *http.Request) {
// 	// flag := make(chan string)
// 	// Извлекаем значение параметра id из URL и попытаемся
// 	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
// 	// конвертировать в integer, или значение меньше 1, возвращаем ответ
// 	// 404 - страница не найдена!
// 	root := r.URL.Query().Get("root")
// 	// if err != nil {
// 	// 	http.NotFound(w, r)
// 	// flag <- root
// 	// }
// 	// fmt.Println(root)
// 	fmt.Sprintf(root)
// 	// Используем функцию fmt.Fprintf() для вставки значения из id в строку ответа
// 	// и записываем его в http.ResponseWriter.
// 	// fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
// }
