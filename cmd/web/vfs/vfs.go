package vfs

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type VFSNode struct {
	Path string  `json:"path"`
	Size float64 `json:"size"`
}

type MainVFS struct {
	Duration *time.Duration `json:"time"`
	Root     string         `json:"root"`
	Node     []VFSNode      `json:"VFSNode struct"`
}

//Терминальная утилита RBS-EX2.3 используется для анализа размера содержимого для указанной директории.

func DirLook(root string) ([]VFSNode, error) {
	start := time.Now()
	// root = flag.String("root", "", "path")
	// flag.Parse()
	duration := time.Since(start)
	var filesOfDir []string
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Ошибка чтения директории %e", err)
	}
	path, err := filepath.Abs(root)
	if err != nil {
		fmt.Printf("Ошибка назначенного пути %e", err)
	}
	filepath.Abs(root)
	for _, file := range files {
		filesOfDir = append(filesOfDir, filepath.Join(path, file.Name()))
	}

	vfsNodes := make([]VFSNode, 0)

	var wg sync.WaitGroup
	wg.Add(len(filesOfDir))
	for _, dirEntered := range filesOfDir {

		go func(dirEntered string) {
			defer wg.Done()
			size := dirSize(dirEntered)
			vfsNodes = append(vfsNodes, VFSNode{Path: dirEntered, Size: size})
		}(dirEntered)
	}
	wg.Wait()

	fmt.Println(duration)
	return vfsNodes, nil
}

/*dirSize() функция принимает путь к директории, определяет тип содержимого (файл или папка)
и возвращает размер содержимого для файла либо сумму размеров содержимого для папки*/
func dirSize(path string) float64 {
	sizes := make(chan int64)
	// booler := make(chan bool)
	size := int64(0)
	readSize := func(path string, file os.FileInfo, err error) error {
		if err != nil || file == nil {
			return err
		}
		if !file.IsDir() {
			// fmt.Println(path, file.Size())
			sizes <- file.Size()
		} else {
			// if file.IsDir() {
			size = file.Size()
		}
		return err
	}
	go func() {
		filepath.Walk(path, readSize)
		close(sizes)
	}()
	for s := range sizes {
		size += s
	}
	return float64(size)

}
