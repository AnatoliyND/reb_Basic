/*
1.	Перейдите в проект с заданиями второго модуля module02 и создайте там новую ветку 04_06_task из ветки 08_task.
2.	В последнем задании 2-го модуля нужно было обработать ошибку io.EOF теперь попробуйте обработать ее с помощью пакета errors.
	- Основная логика должна быть вынесена в отдельную функцию, которая
возвращает результат и ошибку. Это нужно для того чтобы в случае ошибки мы могли проверить ее тип.
3.	Создайте переменную limit укажите ее в произвольное значение и если counter строк превысит лимит вы должны выбросить ошибку.
4.	Ошибка при превышении лимита, должна быть структурой, которая имплементирует интерфейс error и содержит необходимые свойства для вывода: fmt.Sprintf("%s, limit: %d, last string: %s", e.message, e.limit, e.lastString)
5.	В функции main, обработайте кейс с превышением лимита, при помощи пакета errors. Кейс не должен создавать панику, лишь печатать сообщение с ошибкой fmt.Println("string count exceed limit, please read another file =) err: ", err.Error()) в случае, если получена наша имплементация интерфейса error.

*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type errNotFound struct {
	message string
}

type errFinish struct {
	*errNotFound
	limit      int
	lastString string
}

var limit = 9
var laststr string
var errStringLimit = "string count exceed limit, please read another file =) err:"

func (e *errNotFound) Error() string {
	return e.message
}

func readFile(limit int) ([]string, error) {
	//	pathReadFile := "02-08/08_task_in.txt"
	pathReadFile := "08_task_in.txt"

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
		if len(lines) == limit {
			laststr = line
			ee := &errFinish{
				&errNotFound{message: errStringLimit},
				limit,
				line,
			}

			log.Println(fmt.Sprintf("%s, limit: %d, last string: %s", ee.message, ee.limit, ee.lastString))

		}
		if err == io.EOF {
			break
		}
	}
	//fmt.Print(lines)
	//fmt.Printf("\nTotal strings: %d", len(lines))

	return lines, err
}

func main() {

	readFile(55)

}
