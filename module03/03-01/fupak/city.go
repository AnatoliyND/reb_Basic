/* Создайте в проекте moduleOЗ пакет (с произвольным названием), который будет состоять из двух go файлов и содержать 2 функции, по одной в каждом файле:
•	City() - возвращает случайный город,
•	Digit() - возвращает случайное число строчного типа (one, two, three и т.д.).
•	Обе функции должны формировать результат с помощью функции Random из пакета wordz. При этом не внося никаких изменений в пакет wordz.
•	Выведите через fmt.Println результат функции City и Digit в файле main.go
	С помощью утилиты go get, установите пакет для работы со строками xstrings.
	В файле main.go примените функцию Shuffle из этого пакета к результату функции City(). И угадайте, какое название города вывелось :)
*/
package fupak

import "03-01/wordz"

func RandomCity() string {
	WordNew := []string{"Krasnodar", "Moscow", "Omsk", "St. Petersburg", "Anapa", "Sochi", "Innopolis", "Ufa", "Yalta", "Doneck"}
	Wordsss := wordz.Words
	for i := range Wordsss {
		Wordsss[i] = WordNew[i]
	}
	return wordz.Random()
}
