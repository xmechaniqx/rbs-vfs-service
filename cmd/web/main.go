package main

import (
	"log"
	"net/http"
)

// func main() {
// 	mux := http.NewServeMux()

// 	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
// 	mux.Handle("/static", http.NotFoundHandler())
// 	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

// 	err := http.ListenAndServe(":4000", mux)
// 	log.Fatal(err)
// }

// type neuteredFileSystem struct {
// 	fs http.FileSystem
// }

// func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
// 	f, err := nfs.fs.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err := f.Stat()
// 	if s.IsDir() {
// 		index := filepath.Join(path, "index.html")
// 		if _, err := nfs.fs.Open(index); err != nil {
// 			closeErr := f.Close()
// 			if closeErr != nil {
// 				return nil, closeErr
// 			}

// 			return nil, err
// 		}
// 	}

// 	return f, nil
// }

func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// mux.Handle("/ui/static/img/", http.StripPrefix("/ui/static/img/", http.FileServer(http.Dir("./img"))))
	// mux.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "base.layout.html")
	// })

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера.
	// Мы передаем два параметра: TCP-адрес сети для прослушивания (в данном случае это "localhost:4000")
	// и созданный рутер. Если вызов http.ListenAndServe() возвращает ошибку
	// мы используем функцию log.Fatal() для логирования ошибок. Обратите внимание
	// что любая ошибка, возвращаемая от http.ListenAndServe(), всегда non-nil.
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// fmt.Println(fileServer)
	// fileServer := http.FileServer(http.Dir("C:/workspace"))
	// mux.Handle("/#/", http.StripPrefix("/#", fileServer))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
