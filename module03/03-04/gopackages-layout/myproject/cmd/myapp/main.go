/* 1.	Ознакомьтесь со структурой проекта project-layout.
2.	Создайте новый проект(модуль go mod) с названием gopackages-layout.
3.	Код из второго задания модуля перенесите в этот проект, в качестве структуры используйте project-layout.
4.	Используйте только нужные директории.
5.	Запустите его и проверьте, что все работает корректно.
*/
package main

import (
	"fmt"

	"myproject/pkg/fupak"
	"myproject/pkg/wordz"

	"github.com/huandu/xstrings"
)

func main() {

	for i := 0; i < len(wordz.Words); i++ {
		fmt.Println(xstrings.Shuffle(fupak.RandomCity()))
		fmt.Println(fupak.RandomDigit())
	}
}
