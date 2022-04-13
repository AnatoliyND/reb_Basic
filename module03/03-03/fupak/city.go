package fupak

import "03-03/wordz"

func RandomCity() string {
	wordz.Words = []string{"Krasnodar", "Moscow", "Omsk", "St. Petersburg", "Anapa", "Sochi", "Innopolis", "Ufa", "Yalta", "Doneck"}

	return wordz.Random()

}
