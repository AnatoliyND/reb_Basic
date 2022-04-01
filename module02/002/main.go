package main

import (
	"fmt"
	"math"
)

type (
	AmericanVelocy float64
	EuropeanVelocy float64
)

var (
	p  = 1.609
	v1 = 120.4
	v2 = 130
	vA AmericanVelocy
	vE EuropeanVelocy
)

func main() {
	mToKm := p / 3600
	vA = AmericanVelocy(v2) * AmericanVelocy(mToKm)
	fmt.Println(vA)
	fmt.Println("Округляем до 2 знаков: ", math.Round((float64(vA))*100)/100)
	vE = EuropeanVelocy(v1) * EuropeanVelocy(mToKm)
	fmt.Println(vE)
	fmt.Println("Округляем до 2 знаков: ", math.Round((float64(vE))*100)/100)
}
