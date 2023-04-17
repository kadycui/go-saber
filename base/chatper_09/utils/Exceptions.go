package utils

import (
	"errors"
	"fmt"
)

func div(dividend int, divisor int) (int, error) {
	if dividend == 0 {
		return 0, errors.New("divisor is zero")
	}
	return dividend / divisor, nil
}

func Demo7() {
	resl, err := div(4, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("4/2 = ", resl)
	}
	res2, err := div(1, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("1/0 = ", res2)
	}
}
