package fupak

import "03-03/wordz"

func RandomDigit() string {
	wordz.Words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	return wordz.Random()

}
