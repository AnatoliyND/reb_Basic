package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

func startTransaction(debtor internal.Debtor) error {
	return debtor.CalcDiscount()
}

func startTransactionDynamic(d interface{}) error {
	discount, ok := d.(internal.Debtor)
	if !ok {
		return errors.New("Incorrect type")
	}
	err := discount.CalcDiscount()
	return err
}

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	//cust2 := internal.NewCustomer("Dmitry", 23, 10000, 1000, false)

	price := 750

	//cust.CalcDiscount()
	//startTransaction(cust)

	startTransactionDynamic(cust)

	fmt.Printf("%+v\n", cust)
	fmt.Println(internal.CalcPrice(*cust, price))
}
