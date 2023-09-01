package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rbs-vfs-service/cmd/web/vfs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/flag", responseHandler)
	mux.HandleFunc("/stat", stat)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	//Определение файлов необходимых для работы сервера
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
func responseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	root := r.URL.Query().Get("root")
	// flag?root=/home/username/node_modules/
	defaulPath := "/home/username/node_modules/"
	if root == "/" {
		root = defaulPath
		fmt.Println("defaulRoot")
	}
	//Статус чтения из базы данных для заданной пользователем директории
	fmt.Printf("%s", root)
	returner, err := vfs.DirLook(root)
	if err != nil {
		fmt.Println("Ошибка функции vfs.DirLook")
	}
	//Конфигурация ответа в формате JSON
	output, err := json.MarshalIndent(returner, "", "\t")
	if err != nil {
		fmt.Println("Can't Marshall JSON")
	}
	//Отправка ответа серверу APACHE
	phpResp, err := http.Post("http://192.168.81.46/put_stat.php", "application/json", bytes.NewBuffer(output))
	if err != nil {
		fmt.Printf("error of send stat data to php app: %v", err)
	}
	content, err := ioutil.ReadAll(phpResp.Body)
	if err != nil {
		fmt.Printf("error of read response: %v", err)
	}
	fmt.Println("readed response", string(content))
	//Отправка ответа
	w.Write(output)
}

//Хэндлер перенаправления на страницу статистики из базы данных
func stat(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://192.168.81.46/read_stat.php", http.StatusSeeOther)
}
