/*
	Скачайте файл in.txt и скопируйте его в директорию data рядом с файлом main.go.
	Напишите программу, которая построчно считывает файл.
	Выведите в консоль количество строк в формате Total strings: %d.
	Корректно обработайте в отложенном вызове ошибки закрытия файловых дескрипторов.
	Корректно обработайте ошибку окончания файла EOF.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	pathReadFile := "02-08/08_task_in.txt"
	file, err := os.Open(pathReadFile)
	if err != nil {
		err := fmt.Errorf("Cann`t open file: %v", pathReadFile)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lines []string
	for {
		line, err := reader.ReadString(0x0A)
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
	}
	fmt.Print(lines)
	fmt.Printf("\nTotal strings: %d", len(lines))
}
