package main

import (
	fucity "03-01/fupak"
	//fudigit "03-01/fupak"
	"03-01/wordz"
	"fmt"

	"github.com/fatih/color"
	"github.com/huandu/xstrings"
)

func main() {
	color.Red("Hello world again")
	for i := 0; i < len(wordz.Words); i++ {
		fmt.Println(xstrings.Shuffle(fucity.RandomCity()))
		//fmt.Println(fudigit.RandomDigit())
	}
}
