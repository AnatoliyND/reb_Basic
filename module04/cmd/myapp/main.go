package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}

func startTransactionDynamic(d interface{}) error {
	discount, ok := d.(internal.Debtor)
	if !ok {
		return errors.New("Incorrect type")
	}
	err := discount.WrOffDebt()
	return err
}

func main() {

	cust := internal.NewCustomer("Dmitry", 23 /* 10000, 1000, */, true)

	custDisc := &internal.Overduer{
		Balans: 10000,
		Debt:   1000,
	}

	cust.Overduer = custDisc

	price := 750

	//cust.CalcDiscount()
	//startTransaction(cust)

	startTransactionDynamic(cust)

	fmt.Printf("%+v\n", cust)
	fmt.Println(internal.CalcPrice(*cust, price))
}
