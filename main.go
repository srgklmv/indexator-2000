package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	file = flag.String("file", "", "path to file for indexation")
)

func main() {
	fmt.Printf("Индексатор-2000 загружается.\n")

	flag.Parse()

	filename := *file

	if filename == "" {
		fmt.Printf("Ты чё, еблан? Пошёл нахуй с таким файлом.\n") // todo: добавить сюда генерацию файла с медведем с хуем
		return
	}

	_, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("error os.OpenFile: %v\n", err)
	}
}
