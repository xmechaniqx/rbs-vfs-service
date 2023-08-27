package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"snippetbox/cmd/web/vfs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/flag", showFlag)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
func showFlag(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	root := r.URL.Query().Get("root")
	fmt.Println(root)
	returner, err := vfs.DirLook(root)
	if err != nil {
		fmt.Println("Ошибка функции vfs.DirLook")
	}
	// output, err := json.Marshal(returner)
	output, err := json.MarshalIndent(returner, "", "\t")
	if err != nil {
		fmt.Println("Can't Marshall JSON")
	}
	w.Write(output)
}

// -dir="/home/username/workspace/rbs-ex1"
// -root="/home/username/Загрузки"
