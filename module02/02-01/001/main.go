/* Задание:
объявите две переменные, первая - строка со значением 104, вторая -целое число со значением 35;
приведите строку к целому числу, а целое число - к строке;
*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	st string = "104"
	dg int    = 35
)

func main() {
	fmt.Println(st, reflect.TypeOf(st))
	st, _ := strconv.Atoi(st)
	fmt.Println(st, reflect.TypeOf(st))

	fmt.Println(st, reflect.TypeOf(dg))
	dg := strconv.Itoa(dg)
	fmt.Println(dg, reflect.TypeOf(dg))
}
