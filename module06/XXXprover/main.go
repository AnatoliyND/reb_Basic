package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(ReverseInt(-2147483641))
}

func ReverseInt(x interface{}) (int, error) {
	y, ok := x.(int)
	if !ok {
		return 0, errors.New("not int")
	}

	var result int
	for y != 0 {
		result = result*10 + y%10
		if result > 2147483647 || result < -2147483648 {
			return 0, nil
		}

		y /= 10
	}

	return result, nil
}
