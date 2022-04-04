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
