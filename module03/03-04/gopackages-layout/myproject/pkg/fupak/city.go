package fupak

import "myproject/pkg/wordz"

func RandomCity() string {
	WordNew := []string{"Krasnodar", "Moscow", "Omsk", "St. Petersburg", "Anapa", "Sochi", "Innopolis", "Ufa", "Yalta", "Doneck"}
	Wordsss := wordz.Words
	for i := range Wordsss {
		Wordsss[i] = WordNew[i]
	}
	return wordz.Random()
}
