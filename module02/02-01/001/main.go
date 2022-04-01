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
