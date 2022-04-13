/* Создание модулей. Написание собственных библиотек
-	МАЖОРНАЯ версия - когда сделаны обратно несовместимые изменения.
-	МИНОРНАЯ версия - когда вы добавляете новую функциональность, не нарушая обратной совместимости.
-	ПАТЧ-версия - когда вы делаете обратно совместимые исправления.
*/
package main

import (
	"03-03/fupak"
	"03-03/wordz"

	"fmt"

	"github.com/AnatoliyND/utils"
	utilsV2 "github.com/AnatoliyND/utils/v2"
	"github.com/fatih/color"
	"github.com/huandu/xstrings"
)

func main() {
	isExistV2 := utilsV2.InSlice(wordz.Words, "Two")
	if isExistV2 {
		fmt.Println("Slice Int contain finding value")
		return
	}

	isExistInt := utils.ContainsInt([]int{1, 2, 3, 4, 5}, 5)
	if isExistInt {
		fmt.Println("Slice Int contain finding value")
		return
	}

	isExist := utils.Contains(wordz.Words, "Two")
	if isExist {
		fmt.Println("Slice Words contain finding value")
		return
	}
	color.Red("Hello world again")
	fmt.Println(wordz.Hello)
	wordz.Words = []string{"Moscow", "New York", "Amstrdam", "Barselona", "Paris", "London", "Sochi", "Omsk", "Ufa", "Chu", "Petropavlovsk"}
	fmt.Println(wordz.Random())
	wordz.Prefix = ""
	fmt.Println(fupak.RandomCity())
	fmt.Println(fupak.RandomDigit())

	fmt.Println(xstrings.Shuffle(fupak.RandomCity()))
	fmt.Println(xstrings.Shuffle(fupak.RandomDigit()))
}
