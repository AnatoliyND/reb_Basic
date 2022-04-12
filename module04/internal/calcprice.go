package internal

import (
	"fmt"
)

func CalcPrice(z Customer, price int) (int, error) {
	var DiscountedPrice int
	discount, err := z.CalcDiscount()
	if err != nil {
		discount = 0
		fmt.Println(err.Error())
	}
	DiscountedPrice = price - discount
	if DiscountedPrice < 1 {
		DiscountedPrice = 1
		fmt.Println("Discounted price: ")
		return DiscountedPrice, nil
	}
	fmt.Println("Discounted price: ")
	return DiscountedPrice, nil

}
