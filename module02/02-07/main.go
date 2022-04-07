/* Задание:
3.	Скачайте файл in.txt и скопируйте его в директорию data рядом с файлом main.go.
4.	Напишите программу, которая считывает данные из файла, проверяет, что все поля заполнены и записывает считанные данные в файл data/data_out.txt в формате Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n.
5.	Если какое-то поле не заполнено, то программа должна вызвать панику, передав строку вида parse error: empty field оп string %d.
6.	Объявите необходимые отложенные вызовы.
7.	Обработайте панику таким образом, чтобы после обработки на экран вывелось содержимое файла data_out.txt, которое было записано до возникновения паники.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	pathR := "02-07/07_task_in.txt"
	pathW := "02-07/data_out.txt"
	lines, err := readLines(pathR)
	if err != nil {
		log.Fatalf("readLines : %s", err)
	}

	for i, item := range lines {
		arr := strings.Split(item, "|")
		checkingAnEmptyField(i, arr)
		fmt.Printf("%s %d %s %s %s", "Row: ", i+1, "\nName: "+arr[0], "\nAddress: "+arr[1], "\nCity: "+arr[2]+"\n\n\n")
		result := fmt.Sprintf("%s %d %s %s %s", "Row: ", i+1, "\nName: "+arr[0], "\nAddress: "+arr[1], "\nCity: "+arr[2]+"\n\n\n")
		file, err := os.OpenFile(pathW, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString(result)
		file.Close()
	}

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func checkingAnEmptyField(i int, arr []string) (err error) {
	for _, item := range arr {
		if len(item) < 1 {

			//defer panic("")
			//defer fmt.Printf("\nparse error: empty field on string %d: \n", i+1)
			panic(fmt.Sprintf("parse error: empty field on string %d ", i+1))
		}
	}
	return err
}
