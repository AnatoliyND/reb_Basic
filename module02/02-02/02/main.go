package main

import (
	"fmt"
	"math"
)

var lC = 35 //l=2*Pi*r/ S=Pi*r*r

func main() {

	r := math.Round((float64(lC)/(2*math.Phi))*100) / 100
	//r := (float64(lC) / (2 * math.Phi))
	fmt.Println("Радиус окружности = ", r)
	var R *float64
	R = &r
	S := math.Round((math.Phi*math.Pow(*R, 2))*100) / 100
	//S := math.Phi * math.Pow(*R, 2)
	fmt.Println("Площадь круга = ", S)
}
