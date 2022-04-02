package main

import "fmt"

var (
	A *int //объявляем переменнкю A как указатель на int
	B int  = 95
)

func main() {
	A = &B          //присваеваем в переменную A указатель на целочисленную переменную B
	fmt.Println(*A) //выводим разименование указателя A
	*A = 45         // через указатель A присваиваем новое значение целочисленной переменной B
	fmt.Println(B)
}