package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	price := 750

	cust.ResDiscount()
	fmt.Printf("%+v\n", cust)
	fmt.Println(internal.CalcPrice(*cust, price))
}
