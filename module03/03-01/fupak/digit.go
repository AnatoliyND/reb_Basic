package fupak

import "03-01/wordz"

func RandomDigit() string {
	WordNew := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	Wordsss := wordz.Words
	for i := range Wordsss {
		Wordsss[i] = WordNew[i]
	}
	return wordz.Random()
}
