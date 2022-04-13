package main

import (
	"fmt"
	"myapp/internal"
)

func startDiscounter(discounter internal.Discounter) (int, error) {
	return discounter.CalcDiscount()
}

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	price := 750
  
	//cust.CalcDiscount()
	startDiscounter(cust)
  
	fmt.Printf("%+v\n", cust)
	fmt.Println(internal.CalcPrice(*cust, price))
}
