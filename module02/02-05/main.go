/* Задание:
3.	Напишите функцию contains, которая принимает на вход два параметра: слайс строк а и строку х. Функция должна проверять, содержится ли строка х в слайсе а, и возвращать булево значение.
4.	Создайте вариативную функцию getMax, которая находит максимальное целое число из переданных на вход параметров.
5.	Выведите на экран результат вызова
*/
package main

import "fmt"

func main() {
	fmt.Println(contains([]string{"Anna", "Ben", "John", "Den"}, "Den"))
	fmt.Println(getMax(-333, -6, 3, -12, -54, 9))
}

func contains(a []string, x string) bool { //функция для проверки на содержание строки "x" в слайсе "a", на выходе получаем булево значение
	i := ""
	for _, i := range a {
		switch x {
		case i:
			return i == x
		default:
			continue
		}
	}
	return i == x
}

func getMax(x ...int) (maxX int) { // функция для нахождения максимального целого числа из переданных чисел
	maxX = x[0]
	for _, i := range x {
		if i > maxX {
			maxX = i
		}
	}
	return maxX
}
