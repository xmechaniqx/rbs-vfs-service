package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	mux.HandleFunc("/stat", stat)

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

	phpResp, err := http.Post("http://192.168.81.46/put_stat.php", "application/json", bytes.NewBuffer(output))
	if err != nil {
		fmt.Printf("error of send stat data to php app: %v", err)
	}
	content, err := ioutil.ReadAll(phpResp.Body)
	if err != nil {
		fmt.Printf("error of read response: %v", err)
	}
	fmt.Println("readed response", string(content))
	// fmt.Println(resp)
	w.Write(output)
}
func stat(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://192.168.81.46/read_stat.php", http.StatusSeeOther)
}

// -dir="/home/username/workspace/rbs-ex1"
// -root="/home/username/Загрузки"
