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
	Stat string  `json:"stat"`
}

type MainVFS struct {
	Duration time.Duration `json:"time"`
	Root     string        `json:"root"`
	Node     []VFSNode     `json:"VFSNode_struct"`
}

//Терминальная утилита RBS-EX2.3 используется для анализа размера содержимого для указанной директории.

func DirLook(root string) (MainVFS, error) {
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
	for _, file := range files {
		filesOfDir = append(filesOfDir, filepath.Join(path, file.Name()))
	}
	vfsNodes := make([]VFSNode, 0)
	var wg sync.WaitGroup
	wg.Add(len(filesOfDir))
	for _, dirEntered := range filesOfDir {
		go func(dirEntered string) {
			defer wg.Done()
			fmt.Println("dirEntered=", dirEntered)
			var dirEnteredType string
			var size float64
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

			vfsNodes = append(vfsNodes, VFSNode{Path: dirEntered, Size: size, Stat: dirEnteredType})

		}(dirEntered)
	}
	wg.Wait()

	MyMainVFS := MainVFS{
		Duration: duration,
		Root:     root,
		Node:     vfsNodes,
	}
	return MyMainVFS, nil
}

/*dirSize() функция принимает путь к директории, определяет тип содержимого (файл или папка)
и возвращает размер содержимого для файла либо сумму размеров содержимого для папки*/
func dirSize(root string) int64 {
	fmt.Println("dirsize root = ", root)
	var size int64 = 0
	readSize := func(path string, file os.FileInfo, err error) error {
		if err != nil || file == nil {
			return err
		}
		if !file.IsDir() {
			size += file.Size()
		} else {
			size += dirSize(filepath.Join(root, file.Name()))
		}
		return err
	}
	err := filepath.Walk(root, readSize)
	if err != nil {
		fmt.Println(err, "Ошибка filepath.Walk", "path", root, "readSize", readSize)
	}
	return size
}

// http://127.0.0.1:4000/flag?root=/home/username/
