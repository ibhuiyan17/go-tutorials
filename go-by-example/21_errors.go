package main

import (
	"errors"
	"fmt"
)

// multiply func that can't square
func mult_not_square(a, b int) (int, error) {
	if a == b {
		return -1, errors.New("sorry, can't use this func to square")
	}

	return a * b, nil
}

func main() {
	if mult, er1 := mult_not_square(2, 3); er1 == nil {
		fmt.Println(mult)
	} else {
		fmt.Println(er1)
	}

	if square, er2 := mult_not_square(3, 3); er2 == nil {
		fmt.Println(square)
	} else {
		fmt.Println(er2)
	}

}