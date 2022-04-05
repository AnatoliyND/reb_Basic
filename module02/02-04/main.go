/* Задание:
•	Создайте map, в которой необходимо хранить информацию о выданных читателю печатных изданиях: книгах и периодических изданиях.
•	Тип ключей отображения является строкой, тип значений - отображением с
􀀁
ключами типа строка и значениями с типом слаис-строк.
*/
package main

import (
	"fmt"
)

func main() {
	readerInfo := map[string]map[string][]string{}
	/* readerInfo["Anna"] = map[string][]string{}
	readerInfo["Anna"]["MathBook"] = []string{"book1", "book2"} */
	readerInfo["Anna"] = map[string][]string{
		"Books on mathematics": {"book1", "book2", "magazin1"},
		"Books on geometry":    {"book1", "magazin1"},
		"History books":        {"book1", "magazin1", "magazin2"},
	}
	readerInfo["Ben"] = map[string][]string{
		"Books on mathematics": {"book1", "book2", "magazin1"},
		"Books on geometry":    {"book1", "book2", "book3", "magazin1"},
		"History books":        {"book1", "magazin1", "magazin2"},
	}
	readerInfo["Duke"] = map[string][]string{
		"Books on mathematics": {"book1", "magazin1"},
		"Books on geometry":    {"book1", "book2", "book3", "magazin1", "magazin2", "magazin3"},
		"History books":        {"book1", "book2", "magazin1", "magazin2"},
	}

	//fmt.Println(readerInfo)
	fmt.Println("\nКоличество читателей: ", len(readerInfo))

	for k, v := range readerInfo {
		fmt.Println("\nКоличество изданий у", k, ":")
		z := 0
		for i, n := range v {
			fmt.Println(i, "-", len(n))
			if len(n) != 0 {
				z += len(n)
			}
		}
		fmt.Println("Общее количество изданий у", k, "-", z)
	}
}
