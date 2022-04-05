/* Задание:
•	найдите радиус окружности, если её длина равна 35.
•	радиус окружности R объявите как указатель на float64.
•	вычислите площадь круга, используя при расчёте разыменовывание указателя R.
•	При необходимости дробные значения округлите до двух знаков после запятой.
*/
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
