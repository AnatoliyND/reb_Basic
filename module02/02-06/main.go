/* Задание:
Скачайте файл in.txt и положите его в директорию data рядом с файлом main.go.
В файле main.go напишите код, который считывает данные из файла in.txt и построчно записывает их в файл out.txt, нумеруя каждую строку. Если файла out.txt нет, то он должен создаваться.
С помощью отложенных вызовов закройте файловые дескрипторы. При закрытии файла out.txt программа должна вывести в консоль, сколько строк и байт было записано в файл.
Реализуйте функцию logTime(), которая не принимает на вход параметров, определяет и выводит на экран время выполнения вашей программы. Вывод времени должен происходить перед завершением работы функции main().
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var Start time.Time

func main() {
	Start = time.Now()
	defer logTime()
	path := "02-06/06_task_in.txt"
	writeFile := "02-06/out.txt"
	readFile, _ := os.Open(path) //открыть файл
	defer readFile.Close()       //закрыть файл (отложено)
	rf := bufio.NewScanner(readFile)
	var d []string
	var b []byte
	for rf.Scan() {
		d = append(d, rf.Text())
		b1 := []byte(rf.Text() + "\n")
		b = append(b, b1...)
		//fmt.Println(rf.Text())
	}
	//fmt.Println(d)

	cf, _ := os.Create(writeFile) //создать файл

	var z string
	for i, v := range d {
		i += 1
		fmt.Fprintln(cf, i, v)

		z += fmt.Sprintln(v)
	}
	defer func() {
		fmt.Printf("Записано %d строк (%d байт)", strings.Count(z, "\n"), len(b))
		cf.Close()
	}()

}

func logTime() {
	duration := time.Since(Start)
	defer func() {
		fmt.Println("\n Время выполнения программы: ", duration)
	}()

}
