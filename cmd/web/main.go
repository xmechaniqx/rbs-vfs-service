package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"snippetbox/cmd/web/vfs"
)

func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/flag", showFlag)
	// mux.Handle("/ui/static/img/", http.StripPrefix("/ui/static/img/", http.FileServer(http.Dir("./img"))))
	// mux.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "base.layout.html")
	// })

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// fmt.Println(fileServer)
	// fileServer := http.FileServer(http.Dir("C:/workspace"))
	// mux.Handle("/#/", http.StripPrefix("/#", fileServer))
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	// vfs.DirLook(&"s")
}
func showFlag(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	root := r.URL.Query().Get("root")
	fmt.Println(root)
	returner, err := vfs.DirLook(root)
	if err != nil {
		fmt.Println("Ошибка функции vfs.DirLook")
	}
	// fmt.Println(returner)
	// json.NewEncoder(w).Encode(r)
	output, err := json.MarshalIndent(returner, "", "\t")
	if err != nil {
		fmt.Println("Marshall")
	}

	w.Write(output)
}

// -dir="/home/username/workspace/rbs-ex1"
// -root="/home/username/Загрузки"
