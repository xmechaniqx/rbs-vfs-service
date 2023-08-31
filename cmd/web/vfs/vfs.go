//Терминальная утилита RBS-EX2.3 используется для анализа размера содержимого для указанной директории
package vfs

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//Структура VFSNode содержит параметры расположения, размера и типа (файл или папка) вложенных объектов внутри родительской директории (MainVFS.root)
type VFSNode struct {
	Path string  `json:"path"` //Расположение обрабатываемого файла (не путать с корневым)
	Size float64 `json:"size"` //Размер обрабатываемого файла (не путать с корневым)
	Stat string  `json:"stat"` //Статус файлом или директорией является обрабатываемый файл по заданному расположению
}

//Структура MainVFS содержит параметры вычисляемые в процессе выполнения функции DirLook().
type MainVFS struct {
	Duration time.Duration `json:"time"`          //Время затраченное на обработку всех файлов в заданной директории
	Root     string        `json:"root"`          //Расположение корневой директории
	Node     []VFSNode     `json:"VFSNodeStruct"` //Структура описывающая параметры конкретных файлов
	Data     string        `json:"data"`          //Дата и время совершения вычислений
	MainSize float64       `json:"mainsize"`      //Общий размер директории
}

var j = 3672
var i = 0

/*dirLook() функция принимает путь к директории, совершает проход пофайловый обход собирая набор параметров в структуру MainVFS */
func DirLook(root string) (MainVFS, error) {

	start := time.Now()
	var filesOfDir []string //Массив включающий все директории для "root"
	var wg sync.WaitGroup
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Ошибка чтения директории %e", err)
	}
	path, err := filepath.Abs(root)
	if err != nil {
		fmt.Printf("Ошибка назначенного пути %e", err)
	}
	for _, file := range files {
		filesOfDir = append(filesOfDir, filepath.Join(path, file.Name()))
	}
	vfsNodes := make([]VFSNode, 0)
	result := make([]float64, len(filesOfDir))
	wg.Add(len(filesOfDir))
	for i, dirEntered := range filesOfDir { //dirEntered - конкретная директория из массива "filesOfDir[]"
		go func(dirEntered string, i int) {
			defer wg.Done()
			var dirEnteredType string //Переменная задающая тип объекта (файл или папка)
			var size float64          //Итоговый размер объекта
			fs, err := os.Stat(dirEntered)
			if err != nil {
				fmt.Println(err)
			}
			if fs.IsDir() {
				dirEnteredType = "dir"
				size = float64(dirSize(dirEntered))
			} else {
				dirEnteredType = "file"
				size = float64(fs.Size())
			}
			result[i] = size
			vfsNodes = append(vfsNodes, VFSNode{Path: dirEntered, Size: result[i], Stat: dirEnteredType})
		}(dirEntered, i)
	}
	wg.Wait()
	duration := time.Since(start)
	MyMainVFS := MainVFS{
		Duration: duration,
		Root:     root,
		Node:     vfsNodes,
		MainSize: sum(vfsNodes),
		Data:     time.Now().Format(time.RFC850),
	}
	return MyMainVFS, nil
}

//Функция sum() принимает массив структур типа VFSNode и возвращает сумму всех полей параметра "Size"
func sum(vfsNodes []VFSNode) float64 {
	sum := 0.0
	for _, node := range vfsNodes {
		sum += node.Size
	}
	return sum
}

/*dirSize() функция принимает путь к директории, определяет тип содержимого (файл или папка)и возвращает размер содержимого для файла
либо сумму размеров содержимого для папки*/
func dirSize(root string) int64 {
	var size int64 = 0
	readSize := func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !file.IsDir() {
			size += file.Size()
		} else {
			size += dirSize(filepath.Join(root+"/", file.Name()))
		}
		return err
	}
	err := filepath.Walk(root, readSize)
	if err != nil {
		fmt.Println(err, "Ошибка filepath.Walk", "path", root)
	}

	return size
}

// http://127.0.0.1:4000/flag?root=/home/username/
